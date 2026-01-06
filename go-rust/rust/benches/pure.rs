#![feature(test)]
extern crate test;
use test::Bencher;

#[no_mangle]
#[inline(never)]
pub extern "C" fn add(a: i32, b: i32) -> i32 { a + b }

#[bench]
fn bench_add(b: &mut Bencher) {
    b.iter(|| add(1, 2));
}
