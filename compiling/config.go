package compiling

type Compiler string

const (
	Dart       Compiler = "dart"
	Flutter    Compiler = "flutter"
	Golang     Compiler = "go"
	Rust       Compiler = "rust"
	Java       Compiler = "java"
	Maven      Compiler = "maven"
	Nodejs     Compiler = "nodejs"
	TypeScript Compiler = "typescript"
	Python     Compiler = "python"
	Gradle     Compiler = "gradle"
	Dotnet     Compiler = "dotnet"
	GCC        Compiler = "gcc"
	GXX        Compiler = "g++"
	Clang      Compiler = "clang"
	ClangXX    Compiler = "clang++"
	Msvc       Compiler = "msvc"
)

type CompileConfig struct {
	Type Compiler `json:"type"`
}
