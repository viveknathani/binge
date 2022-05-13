# binge

binge is a server to stream content from my library of TV-Shows and Movies. There's a lot of work left here. But it is demo worthy as of now.

## why?

- I store my content on external hard drives. Whenever I need to watch something, I need to copy it into my laptop. But laptop does not have infinite disk space. After a while, the copied content needs to be deleted. It would be nice if I could access my content like we access content on Netflix or Youtube.

## streaming

I am using MPEG-DASH to stream my videos. Once the video is uploaded, it goes to the processor's job queue. Workers of the processor pick up the video from the job queue and encode it as per MPEG-DASH requirements. The .mpd file is generated. When the video is being processed, it can be found in a list of pending jobs seen on the "/jobs" page.

## tech

- Go
- PostgreSQL

## source description

- [database](./database): provides an API to work with the database 
- [entity](./entity): provides the `struct` to represent the entities in the system 
package.
- [processor](./processor/): provides an API for dealing with a data structure that handles video processing for MPEG-DASH.
- [scripts](./scripts): for things like setting up PostgreSQL one time
- [server](./server): implements the routes
- [service](./service): the core logic of this system, to be used by `server`
- [shared](./shared): for things like managing `context` across `server` and `service` 
- [web](./web): the frontend. 


## contributing

- `make build`, `make test`, and `make run` will be some handy commands for you. 
- It would be cool if you write an issue before working on a PR. 

## license

[MIT](./LICENSE)

