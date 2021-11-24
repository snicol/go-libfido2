package libfido2

/*
#cgo darwin,!arm64 LDFLAGS: -framework CoreFoundation -framework IOKit /usr/local/lib/libfido2.a /usr/local/opt/openssl@1.1/lib/libcrypto.a ${SRCDIR}/darwin/lib/libcbor.x86.a
#cgo darwin,arm64 LDFLAGS: -framework CoreFoundation -framework IOKit /opt/homebrew/lib/libfido2.a /opt/homebrew/opt/openssl@1.1/lib/libcrypto.a ${SRCDIR}/darwin/lib/libcbor.arm.a
#cgo darwin,!arm64 CFLAGS: -I/usr/local/opt/libfido2/include -I/usr/local/opt/openssl@1.1/include
#cgo darwin,arm64 CFLAGS: -I/opt/homebrew/opt/libfido2/include -I/opt/homebrew/opt/openssl@1.1/include
*/
import "C"
