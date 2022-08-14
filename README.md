### FlatBuffers

This code allows for translating the users from jsonplaceholder to flatbuffers.

The first part of the code, the "write" part, takes the records from jsonplaceholder.typicode.com/people and writes it out, in flatbuffer format, to the file `people.dat`. This structure has a list of Person records.

The second part of the code, the "read" part, uses an mmap file so only the parts that are required are loaded into memory. The command line option refers to the index of the record you want to read.

This allows for the buffers to not take up a lot of memory while doing I/O.
