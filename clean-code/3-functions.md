## SMALL!
- Functions should be

### Blocks and Indenting
- Max 1-2 indent lines
- Attempt to make the body of conditionals one line (a function call)

## DO ONE THING
- Remember that if a function does only those steps that are one level below the stated name of the function, it only does one thing.
- Example: RenderPageWithSetupsAndTearDowns - 
  - This function arguably does 3 things: 
    - determine if the page is a test page
    - if so, indlude setups and teardowns 
    - render the page in html
  - This function still does one thing (see first bullet)

### Sections within Functions
- If you find that your function can be broken into sections, it likely doesn't do one thing.

## ONE LEVEL OF ABSTRACTION PER FUNCTION

### Reading Code from Top to Bottom: The Stepdown Rule

## SWITCH STATEMENTS
- can be tolerated if they appear only once
- are used to create polymorphic objects
- are hidden behind an inheritance relationship

## USE DESCRIPTIVE NAMES
- Half the "clean code" battle is choosing good names for small functions that do one thing.
- A long descriptive name is better...
  - than a short engmatic name
  - than a long descriptive comment
- Be consistent in your names

## FUNCTION ARGUMENTS**
- zero is ideal (niladic)
- next comes one (monadic)
- followed by two (dyadic)
- three args should be avoided where possible
- More than 3, never.

### Flag Arguments - isTrue
- These are ugly and immediately make a function do more than one thing.

## HAVE NO SIDE EFFECTS
- 'nuff said
- in case I forget: A side effect is a thing that function does outside of its stated purpose.

### Output Arguments
- if there is an output, make it explicit

## PREFER EXCEPTIONS TO RETURNING ERROR CODES

### Extract Try/Catch Blocks
- Try/catch blocks are ugly in their own right.
  - confuse the structure of the code
  - mix err processing with normal processing
- For this reason, its better if the boyd of try/catch blocks are function in their own right.  

### Error Handling Is One Thing

## DON’T REPEAT YOURSELF13

## STRUCTURED PROGRAMMING**
- Dijkstra said that every function and every block within a function should have one entry point and one exit. Following these rules means...
  -  that there should only be one return statement in a function
  -  no break/continue starements in a loop

## HOW DO YOU WRITE FUNCTIONS LIKE THIS?
- Typically start with gross
- Refine it with tests
- keep the tests passing!

## CONCLUSION

*"Master programmers think of systems as stories to be told rather than programs to be written. They use the facilities of their chosen programming language to construct a much richer and more expressive language that can be used to tell that story."*