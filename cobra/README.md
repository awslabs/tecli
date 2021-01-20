# cobra

## Packages

Below you can find the responsibility and purpose for each package.

### aid

Assist Cobra commands individually. Example: a command under `cmd/foo.go`, has its respective `aid/foo.go`.
This allows a more cleaner and readable code.

### cmd

Cobra commands. [reference](https://github.com/spf13/cobra)

### controller

> Controller acts on both model and view. It controls the data flow into model object and updates the view whenever data changes. It keeps view and model separate. [reference](https://www.tutorialspoint.com/design_pattern/mvc_pattern.htm)

### dao (Data Access Object)

> Data Access Object Pattern or DAO pattern is used to separate low level data accessing API or operations from high level business services. Following are the participants in Data Access Object Pattern.
> Data Access Object Interface - This interface defines the standard operations to be performed on a model object(s).
> Data Access Object concrete class - This class implements above interface. This class is responsible to get data from a data source which can be database / xml or any other storage mechanism.
> Model Object or Value Object - This object is simple POJO containing get/set methods to store data retrieved using DAO class. [reference](https://www.tutorialspoint.com/design_pattern/data_access_object_pattern.htm)

### model

> Model represents an object carrying data. It can also have logic to update controller if its data changes. [reference](https://www.tutorialspoint.com/design_pattern/mvc_pattern.htm)

### view

> View represents the visualization of the data that model contains. [reference](https://www.tutorialspoint.com/design_pattern/mvc_pattern.htm)


