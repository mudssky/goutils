$version = "$(svu next)"
trap {
    Write-Output "Error: $($_.Exception.Message)"
}
git tag $version
git push --tags