

#Task 1

As a part of internship the first task is to create User Activity Logger, using GRPC and MongoDB

problem statement: User Activity Logger

•Create a system to track users daily activities. Daily activities includes, “play”, “sleep”, “eat” and “read”. Each record will have activity type, time spent and timestamp of the activity creation. We should be able to add new users, update their activity, query the users and activity.

•Should use gRPC

•Activity is an interface, implement Play, Sleep, Eat and Read activities.

•Activity will have isDone(), isValid() methods

•isDone() bool : can be retrivied by activity.status == DONE

•isValid() err : is a custom implementation, for ex: 6h < sleep < 8h is a valid sleep condition. You are free to set any conditions for this. This can also validate all other inputs.

•Protobuf usage is a plus

•Users will have - Name, Email, Phone

•Activity - Type, Timestamp, Duration, Label

•Code standards, quality and documentation is key
task1
