// Exported directly with `no_mangle` and C ABI for Go to call via cgo

#[no_mangle]
#[inline(never)]
pub extern "C" fn add(a: i32, b: i32) -> i32 {
    a + b
}
