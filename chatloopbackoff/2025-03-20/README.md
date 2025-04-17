# Episode 51: WasmEdge  

- HosT: [Faeka Ansari](https://www.linkedin.com/in/faeka/)  
- Date: 2025/03/20
- Videos: [Episode 51: WasmEdge](https://www.youtube.com/live/Cxz7pC9Lq2k)

## News
- WasmEdge 0.14.1 Released recently, some of it's features include
  - Supported LLVM 17.0.6  
  - Added AI backends: whisper.cpp, piper, ChatTTS, Burn.rs  
  - Introduced `wasmedge_stablediffusion` plug-in  
  - Enabled CUBLAS and Metal support on macOS  
  - Initial support for the WebAssembly component model  
  - Implemented the WASM Relaxed-SIMD proposal  

## Summary  

- WasmEdge  
  - CNCF Sandbox Project (Since 2021)  
  - What?  
    - WasmEdge is a lightweight, high-performance WebAssembly runtime optimized for cloud-native, edge computing, and AI workloads.
  - Why?
    - WasmEdge provides a secure, efficient, and cross-platform execution environment for WebAssembly.
    - Ideal for microservices, serverless functions, and embedded AI applications.  
  - Repository: [https://github.com/WasmEdge/WasmEdge](https://github.com/WasmEdge/WasmEdge)  
  - Website: [https://wasmedge.org/](https://wasmedge.org/)  

  - Talks
    - [WasmEdge: Portable and Lightweight Runtime for AI/LLM Workloads](https://youtu.be/NuLFdERpGSY)
    - [WasmEdge: Cross-Platform, High-Performance, Lightweight, Embeddable Multi-Modal LLM Runtime | PLT](https://youtu.be/5_wTuySm7lE)
  - Slack: Join the [#WasmEdge](https://cloud-native.slack.com/archives/C01QDJNSJ7Q) channel on CNCF Slack
  - Founders: Michael Yuan (Co-founder of [Second State](https://secondstate.io/))

## FLOW

- 🔬 Explored in the Talk  
  - ✅ Checked Website & GitHub Repo
  - ✅ Ran 3 Demos:
    - HTTP Server
    - TinyGo WebAssembly
    - Docker + WasmEdge
  - ✅ Explored WasmEdge Execution Flow
  - ✅ Dived Into the Codebase a bit

- 🌟 Interesting Findings  
  - Strong AI workload support with WASI-NN.
  - WebAssembly + Docker has the power to unlock lightweight, secure workloads.
  - WasmEdge plug-ins extend capabilities (example, Stable Diffusion, AI).
