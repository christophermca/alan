# Who/what controls state of self

## Desire

- -1 Repultion
- 0 Neutral
- 1 Desire

ex. Food

**self is full**

Food:  < 0

**self is hungry**

Food: > 0

# Communicate

### interprating speach

Break string down to pieces to build understanding.

#### interprater
A function used to understand conversation:

on initialize:
- determine the `_actor`
  `_actor` = new actor()
  starts object in runtime memory

|legend |
| :--- |
| [item] |
|`<optional item>` |


### Math
"3 + 3 = 6"
"[numbers]<space>[operator]<space>[numbers]<=><answer>"

### Asking Questions
"how are you?"

#### Determine Subject
``` sh
# Context
actor is speaking to alan
alan <~ actor
```

questions  to>> alan
|Question|Subject|Action|Result|
| :--- |:--- |:--- |:--- |
|"what is your name?"|"[yY]ou" \|\| "[yY]our"  == self|lookup|self.name|
|"what is my name?"|"[mM]y \|\| [iI] == actor |lookup| subject.name|
|"what is your favorite color?"| self | lookup | self.lookup([favorite color])|
|"what is my favorite color?"| actor | lookup  | subject.lookup([favorite color])|
|"Do I have a favorite color?"| actor | lookup | subject.lookup([favorite color])|


Proposed Tools:

**Golang**

__WHY__ I never worked in golang before and think it will be an insteresting language to
learn.

I am also intrigued by the idea of multi threading functions.  I think the ideas
of goroutine will be helpful for this project

**GraphQL**

__WHY__
I theorize that GraphQL will a better data format to save the linked threads of
conversation and idea.

