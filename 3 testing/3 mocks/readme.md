# Mock testing
It is necessary when you need to test calls to external APIs,
or simply complex functions of which you only need their results
rather than how they are generated.

# Keys
- Overwrite (at the test execution time) the functions that perform the API calls,
  modifying their behavior to return a mock data.
- Move the original functions to another variables at the beginning of the test.
- Restore the address of the original functions at the end of the test.