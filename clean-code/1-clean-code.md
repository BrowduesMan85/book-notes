## THERE WILL BE CODE
*"The folks who think that code will one day disappear are like mathematicians who hope one day to discover a mathematics that does not have to be formal. They are hoping that one day we will discover a way to create machines that can do what we **want** rather than what we **say**."*
## BAD CODE

- Bad code brings down companies
- LeBlanc’s law: Later equals never.

## THE TOTAL COST OF OWNING A MESS

As mess builds -> Productivity approaches 0 -> As productivity decreases, management hires more staff -> Additional staff don't know the design of the system -> More messes -> Productivity is driven closer to 0

### The Grand Redesign in the Sky

- Eventually the team rebels, demands a re-design 
- Mgmt relents, tiger team is selected 
- Now its a race between the new product and legacy prod. Mgmt will not release new prod 'til it does everything old prod did  
- The race can go on for a very long time. By the time its done the original tiger team is gone, and the new members of the tiger team are demanding a re-design because of the mess the new prod is in.

*"... keeping your code clean is not just cost effective; it’s a matter of professional survival."*

### Attitude

Its our job, as software developers to defend code cleanliness. Its simply not professional to do anything else.

### The Primal Conundrum

- Devs feel the pressure to make messes in order to meet deadlines. 
- *They don't take the time to **go fast**.*
- The only way to meet the deadline - to go fast - is to keep the code as clean as possible at all times

### The Art of Clean Code?

Its no good to *try* and write clean code if you don't know what is means for code to be clean!

### What Is Clean Code?

- Bjarne Stroustrup (inventor of C++)
  - elegant/efficent (pleasingly graceful)
  - straightfoward logic
  - minimal dependencies, to ease maintanence
  - error handling complete according to an articulated strategy
  - **does one thing well**
    - It's focused!

- Grady Booch 
  - simple/direct
  - reads like well-written prose
  - never obscures the designer's intent, but rather is full of crisp abstractions and straightfoward lines of control

- "Big" Dave Thomas (founder of OTI)
  - can be read/enhanced by others
  - has unit/acceptance tests
  - meaningful names
  - provides one way for doing things
  - minimal deps
    - explicityly defined
    - provide clear/minimal API
  - literate (aka: readable by humans)

- Michael Feathers 
  - looks like its written by someone who cares

- Ron Jeffries
  - runs all the tests
  - contains no duplication
  - expresses all design ideas in the system
  - minimizes the number of entities (classes, methods, functions etc)
  
- Ward Cunningham (godfather of those who care about code)
  - each routine you read turns out to be what you expected
    - aka: Its easy to read
  - beautiful when the code makes it look like the language was made for the problem
    - This implies its our responsiblity to make the language look simple


## SCHOOLS OF THOUGHT
- No "school of thought" is absolutely right.
  
## WE ARE AUTHORS
- Since we spend 10:1 readingCode:writingCode, we **must** write clean code!


## THE BOY SCOUT RULE
*"Leave the campground cleaner than when you found it."*

