#!/bin/bash
# publish.sh builds and publishes a release.
set -o nounset -o errexit -o pipefail
ROOT=$(dirname $0)/..

echo "Publishing NPM packages to NPMjs.com:"

# For each package, first create the package.json to publish.  This must be different than the one we use for
# development and testing the SDK, since we use symlinking for those workflows.  Namely, we must promote the SDK
# dependencies from peerDependencies that are resolved via those links, to real installable dependencies.
publish() {
    node $(dirname $0)/promote.js ${@:2} < \
        ${ROOT}/nodejs/$1/bin/package.json > \
        ${ROOT}/nodejs/$1/bin/package.json.publish
    pushd ${ROOT}/nodejs/$1/bin
    mv package.json package.json.dev
    mv package.json.publish package.json

    NPM_TAG="dev"
    PKG_NAME=$(jq -r .name < package.json)
    PKG_VERSION=$(jq -r .version < package.json)

    # If the package doesn't have a pre-release tag, use the tag of latest instead of
    # dev. NPM uses this tag as the default version to add, so we want it to mean
    # the newest released version.
    if [[ "${PKG_VERSION}" != *-* ]]; then
        NPM_TAG="latest"
    fi

    # Now, perform the publish. The logic here is a little goofy because npm provides
    # no way to say "if the package already exists, don't fail" but we want these
    # semantics (so, for example, we can restart builds which may have failed after
    # publishing, or so two builds can run concurrently, which is the case for when we
    # tag master right after pushing a new commit and the push and tag travis jobs both
    # get the same version.
    #
    # We exploit the fact that `npm info <package-name>@<package-version>` has no output
    # when the package does not exist.
    if [ "$(npm info ${PKG_NAME}@${PKG_VERSION})" == "" ]; then
        if ! npm publish -tag "${NPM_TAG}"; then
	    # if we get here, we have a TOCTOU issue, so check again
	    # to see if it published. If it didn't bail out.
	    if [ "$(npm info ${PKG_NAME}@${PKG_VERSION})" == "" ]; then
		echo "NPM publishing failed, aborting"
		exit 1
	    fi

	fi
    fi
    npm info 2>/dev/null

    # And finally restore the original package.json.
    mv package.json package.json.publish
    mv package.json.dev package.json
    popd

    ## Publish PIP Packages
    echo "Publishing Pip package to pypi as pulumi:"
    pushd ${ROOT}/python/bin
    twine upload \
        -u "pulumi" -p "${PYPI_PASSWORD}" \
        "dist/*.tar.gz" \
        --skip-existing \
        --verbose
    popd

    ## Publish Nuget Package if any exists
    find "${ROOT}/dotnet/bin/Debug/" -name 'Pulumi.*.nupkg' \
            -exec dotnet nuget push -k "${NUGET_PUBLISH_KEY}" -s https://api.nuget.org/v3/index.json {} ';'
}

publish wordle
