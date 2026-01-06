# rust2gosample

A minimal demonstration project showing how to bridge **Rust** and **Go** using  
[`rust2go`](https://github.com/ihciah/rust2go) — a tool that automatically generates Go bindings from Rust code.

This repository focuses on **practical Rust → Go interoperability** with minimal boilerplate.

---

## What This Project Does

- Uses **Rust** as the core implementation language
- Exposes Rust APIs to **Go** via generated bindings
- Leverages **rust2go** to avoid manual cgo glue code
- Demonstrates a clean cross-language build workflow

This is intended as a **learning sample**, not a production-ready framework.

---

## Technology Stack

- **Rust** — core logic and exported APIs
- **Go (Golang)** — consumer of Rust functionality
- **rust2go** — automatic binding generator
- **cgo** — Go ↔ native interface layer

---

## Repository Layout

rust2gosample/
├── src/ # Rust source code
├── Cargo.toml # Rust package configuration
├── go/ # Generated Go bindings + Go example
├── build.rs # rust2go integration
├── README.md
└── .gitignore
*(Structure may evolve as the sample grows.)*

---

## How rust2go Fits In

Instead of manually writing C headers and cgo wrappers, **rust2go**:

1. Reads Rust source files
2. Identifies exported functions and types
3. Generates Go-compatible bindings automatically
4. Produces a Go package ready to import

This keeps Rust as the **single source of truth**.

---

## Prerequisites

Ensure the following tools are installed:

- **Rust** (stable toolchain)
- **Go** (1.20+ recommended)
- **cargo**
- A C compiler (`clang` or `gcc`)
- `rust2go` installed:

```bash
cargo install rust2go
