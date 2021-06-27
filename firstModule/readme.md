# Making my first Go module
- Notes:
## Start a module that others use
1. In a module, you collect one or more related packages for a set of slick set of functions.
2. Go code is grouped in modules
3. Modules specify dependencies, including Go version and other modules needed
4. go.mod files track your dependencies
5. In Go, a function whose name starts with a capital letter can be called by a function not in the same package. 
    - This is called an "exported name"
6. := is a shortcut for declaring and initializing a variable in one line
## Call your code from another module
1. Go code executed as an application must be in main package
2. For production use, you'd publish the module from its repo, where go tools could find it to download it.