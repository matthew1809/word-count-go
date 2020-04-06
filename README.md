# Wordcount TDD

Topic: count words in a file passed as argument on the CLI and output info about the file (number of words, ...)

* First, count words ins
    * Get comfortable with Go syntax and tooling
    * Introduce tests
    
* Then, do it concurrently
    * Introduce goroutines, chans + concurrency patterns
    * Aggregation operator + synchronization points
    * Adapt tests

* Then, count frequency for each word
     * Thread safety on datastructures
     * Adapt tests

* Then input from or output to external system (api, db, ...)
    * Interfaces, dependency mocking, dependency injection, integration tests
    * Adapt tests