## directory-size
A command-line utility to list the size of a list of directories and the cumulative total Size

## Usage

`directory-size` takes a list of directory names separated by commas, like `./directory-size "photos,documents"` and lists the sum total size of the files within the directory. Files that cannot be read will be skipped.

By default, `directory-size` considers 1KB to be 1000 bytes. To use 1024 bytes for 1KB, pass in the argument --1024

Build the cli:
`go build`

Check some directory sizes:
`./directory-size "photos,documents"`

If you add the recursive flag `--recursive` the app will walk nested directories and add their sizes to the root directory's total.

If you add the human-readable flag `--human` the app will display directory sizes in a human-readable format (i.e. "3KB" instead of "3000")

If you add the flag `--1024` the human-readable display will use 1024 bytes as 1KB per the JEDEC 100B.01 standard
