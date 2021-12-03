## Many to one and many to many 
- Document model does not support joins very well.
  - application might need to handle the joins which is not a trivial task.
- Modeling many to one and many to many in document model means redundant data or in-application join processing.
- Document data take a hierarchical model and store nested records within parent record instead of a separated table.

Pros and cons of document model
- schema flexibility
- better locality
- closer to some application's needs
  - if an application naturally adopt the document model's nested structure, then the RMDB style will lead to more complicated code.
- you cannot refer to a certain item within a document
- poor support to joins
  - avoid if application need to use many-to-many relationship
- schema flexibility is not schema free
  -  RMDB enforce the schema when user write data 
  -  but document model implicit requires a schema when user read the data.
    - and it is the user's duty to check if the data has an expected schema/structure when reading.
