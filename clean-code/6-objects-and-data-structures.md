## DATA ABSTRACTION

## DATA/OBJECT ANTI-SYMMETRY
- Objects
    - hide their data behind abstractions
    - expose functions that operate on that data
    - makes it easy to add new classes without changing existing functions

- Data Structures (procedural code)
  - types expose their data and have no meaningful functions
  - makes it easy to add new functions without changing the existing data structures
  
  Fundamentally: 
  - If you will add new types (classes) use OO
  - If you will add new functions (types are pretty well fixed), use data structures and procedural programming


## THE LAW OF DEMETER
*"A module should not know about hte innards of the objects it manipulates.*

- This means an object shouldn't expose its internal structure through accessors (get/set) b/c to do so is to expose, rather than to hide its internal structure

More precisely, a method `f` of a class `C` should only call the methods of these:
- `C`
- An object created by `f`
- An object passed as an arg to `f`
- An object held in an instance variable of `C`
  
The method `f` should **not** invoke methods on objects that are returned by any of the allowed functions above!!

-> Talk to friends, not strangers.

### Train Wrecks
`final String outputDir = ctxt.getOptions().getScratchDir().getAbsolutePath();`

VS

`final String outputDir = ctxt.options.scratchDir.absolutePath;`

The former is a train wreck. Remember, objects should not expose their innards! The later is clearly a data structure.

### Hybrids
- Avoid this approach where both data structures and meaningful functions are used.
  - tough to add new functions
  - tough to add new data structures

## CONCLUSION
- Objects expose behavior and hide data
- Data structures expose data and have no significant behavior