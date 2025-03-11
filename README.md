# gc-piler

gc-piler is a project that converts Go language code to C++. The code for this project is written in Go.

## Advantages

- **Performance**: C++ is known for its high performance and efficiency. By converting Go code to C++, you can leverage the performance benefits of C++.
- **Interoperability**: This tool allows Go developers to integrate their code with existing C++ projects, making it easier to use Go in environments where C++ is the standard.
- **Portability**: C++ is widely supported across different platforms and compilers, making the converted code more portable.
- **Optimization**: C++ provides more control over system resources and memory management, allowing for more fine-tuned optimizations.

## Use Cases

- **Legacy Systems**: Integrate modern Go code with legacy C++ systems without rewriting the entire codebase.
- **Performance-Critical Applications**: Convert performance-critical Go code to C++ to take advantage of C++'s efficiency.
- **Cross-Platform Development**: Use Go for development and convert to C++ for deployment on platforms where C++ is more suitable.

## How to Run

1. Add your Go code to `main.go`.
2. Navigate to the `gc-piler` directory:
   ```sh
   cd gc-piler
3. Use ```go2cpp main.go``` to execute your go code and convert your code to c++