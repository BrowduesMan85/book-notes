## USE INTENTION-REVEALING NAMES
- A comparison between poorly named code and well-named code is discussed
  - getThem -> getFlaggedCells
  - theList -> gameBoard
  - etc..
## AVOID DISINFORMATION
- avoid words whose entrenched meanings vary from our intended meaning
  - e.g don't name a variable 'pwd' b/c it has an established meaning in linux

- don't refer to a grouping of accounts as accountLIst unless its **actually a list**
  - Even if the container is a list, its probably better not to encode the container type into the name.

## MAKE MEANINGFUL DISTINCTIONS
- `moneyAmount` is indistinguishable from `money`
- `theMessage` is indistinguishable from `message`
- etc...

## USE PRONOUNCEABLE NAMES
- ... so meaningful discussion/reasoning about this code is possible

## USE SEARCHABLE NAMES
- One can grep for `MAX_CLASSES_PER_STUDENT` but not for `m`
## AVOID ENCODINGS

### Interfaces and Implementations
- Prefer `ShapeFactory` over `IShapeFactory`
  - no need to know `ShapeFactory` is an interface, just that its a shape factory
  - addornents (like `I`) can be distracting

## CLASS NAMES
- Should have a noun or noun phrase
- Should not be a verb
- Avoid words like manager, data, processor, or info

## METHOD NAMES
- should have a verb or verb phrase (`postPayment`, `deletePage`, `save`)

- when constructors are overloaded, use static factory methods with names that describe the args. For example:

  - ```
    // good
    Complex fulcrumPoint = Complex.FromRealNumber(23.0);
    
    // bad
    Complex fulcrumPoint = new Complex(23.0);
    ```

## DON’T BE CUTE
- Choose `DeleteItems` over `HolyHandGrenade`
  
## PICK ONE WORD PER CONCEPT
- e.g: Its confusing to have fetch, retrieve and get as equivelent methods of different classes

## DON’T PUN
- Don't choose a name for consistency sake only. If they are semantically different, use a different word.

## FINAL WORDS
- People are often afraid of renaming things for fear that some other devs will object. Don't buy into this. Be grateful when names change for the better.