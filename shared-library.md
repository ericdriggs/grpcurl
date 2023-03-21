TODO: clean up / improve notes before merge

/* mainExport calls main function using parameter instead of stdin
allows calling main from other languages when use buildmode c-shared
supports all flags

usage:

    pass the same arguments you would pass as if calling from the command-line, only as a parameter array (GoSlice)
    do not pass program name (grpcurl) -- this is automatically populated

osx - shared library installation

    # build as shared library
    git clone https:github.com/fullstorydev/grpcurl.git
    cd grpcurl/cmd/grpcurl
    mkdir -p lib
    go build -buildmode=c-shared -o lib/grpcurl.dylib

    # install
sudo mkdir -p /usr/local/lib/grpcurl/
sudo mv -f lib/grpcurl.dylib /usr/local/lib/grpcurl/grpcurl.dylib
sudo mkdir -p /usr/local/include/grpcurl/
sudo mv -f lib/grpcurl.h /usr/local/include/grpcurl/grpcurl.h

    # verify dylib file
    otool -L /usr/local/lib/grpcurl.dylib

    # verify loaded after rebooting machine
    otool -D grpcurl.dylib


TODO: exported code shows as type void -- ensure able to access exit code
may need to use a function for os.Exit with a flag to throw instead of return on non zero
which mainExport can then catch and return
*/
export2 mainExport
func mainExport(osArgs []string) int {
	args := append([]string{"grpcurl"}, osArgs...)
	shouldExit = false
	os.Args = args
	return doMain()
}

func exitOrReturn(code int) int {
	if shouldExit {
		os.Exit(code)
	}
	return code
}
