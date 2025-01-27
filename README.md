# directory-size
A command-line utility to list the size of a list of directories and the cumulative total Size

# Usage

`directory-size` takes a list of directory names separated by commas, like `./directory-size "photos,documents"`

Build the cli:
`go build`

Check some directory sizes:
`./directory-size "photos,documents"`

If you add the recursive flag `--recursive` the app will walk nested directories and add their sizes to the root directory's total.

If you add the human-readable flag `--human` the app will display directory sizes in a human-readable format (i.e. "3KB" instead of "3000")
