# MorseMe

Hey, welcome to the dumbest project in existence! This is a study project I have come up with which will involve the following languages, tools and technologies:

- Go (w/ Templ + Echo)
- HTMX
- Hyperscript
- SQLite
- gRPC

Many of these are technologies I've wanted to use or learn more about, there will of course also be other libraries and tools needed to glue all of this together and produce an cohesive solution at the other end, but these are the heavy lifters.

I will include credits for everything I use, look for credits.md in the repo!

## So, what am I (planning on) building exactly?

Glad you asked. MorseMe will be the most useless thing in existence, very literally, it'll deliver a website (written in HTMX + Hyperscript) where visitors will be able to submit a message, the backend server (written in Go, and using SQLite and gRPC) will handle and store that message until a client (written in Go, and using gRPC) checks in and takes the message which it will then output as morse code to an LED array and a speaker, that's it.

It's project splits into two parts; the client and the server, but this breaks down to four areas of concern in total that employ specific tools and technologies:

![MorseMe Overview](overview.png "MorseMe Overview")

## How can I use it?

Go to https://morseme.ryankun.moe/.

Or if you want to *use* it, download or clone the repo and look for setup.md, the plan is nothing will be hardcoded, everything that might change should be in a TOML config file somewhere, and setup.md will be your whistle stop tour of where these TOML files live.