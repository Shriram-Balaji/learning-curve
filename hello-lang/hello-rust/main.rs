fn main() {
    println!("Hello Rust!");
}

/* 
 Takeaways:
 - main fn - first code of execution
 - ! (The exclamatory mark denotes a macro in Rust)
 - Notice that println is a macro. Its a macro for a couple of reasons, macros in rust can have formatting specifiers ( the ones like %s) these need to be checked at compile time.
 - Functions have a fixed number of args, macros can have variable args
 - macros work with call by ref. even when passed as values

 Compiling a rust program: ``` $ rustc main.rs ```
*/