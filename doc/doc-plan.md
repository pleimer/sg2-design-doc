Outline

Title: A Plugin Approach to Smart Gateway

Authors: Paul Leimer

Overview: 
This document proposes a new architecture for the Smart Gateway component of the
Service Telemetry Operator (STO). The cloudops team has been maintaining a
poorly designed code base for this component that hinders performance,
maintainance and expandability of the component. Engineers have reported 
suprise migrains when accidently viewing the code base for a breif moment. 
This is not merely a matter of convenience, but one of health.

Context:
What is the problem that needs to be solved? 
What is the problem with the current solution, and what needs to be fixed?

Goals of this Document:
- Present pluggable solution to SG. Goals of this architecture:
    - to enforce DRY software concepts
    - make code base maintainable
    - decouple code logic for message types and databases with core SG
      functionality to make adapting it to new message types and databases as
      simple as possible 
    - keep documentation simpler - changes to plugins can be documented
      separately than changes to the main codebase
    - Standardize logging, plugin config, and AMQP in external library
- Demonstrate benefits of this approach in the context of the cloud
- Give fleshed out interface APIs for input/output plugins
- Appendix of examples how this might be implemented in golang
- Suggest opensource alternatives to creating our own pluggable engine
- Investment cost for making our own solution and using an opensource one
- ROI section
- Milestones along with dates

Existing Implementation:

- Code base hardcoded to work with Prometheus and ElasticSearch. All one
  monolith application 
- Also hardcoded to work with collectd metrics and events and ceilometer metrics
  and events (Hard coded for 4 slightly different use cases)
- User story:
    1) I can specify collectd or ceilometer, events or metrics as one of many
       CLI arguments
- Programmers story: 
    1) A need arises to implement logging. Must use a different database
        - I have to duplicate much of the messaging code between new message
          types
        - I have to implement new database logic in the core. 
        - No clearly defined interfaces or division of work means duplicating
          lost of code from the other message/backend types implemented
        - Recompile the new smart gateway with additional command line arguments
          to allow users to specify this new type
        - Must release a new version of Smart Gateway every time someting new
          message front ends or backends are added, even if it is just a small
          parsing change
        - All in all, code is kinda a mess, hard to upkeep, and hard to expand

New Implementation:

- Core functionality is Smart Gateway - that is, purely a passthrough engine
  that calls plugins for parsing, formatting, storage and caching of messages.
  Actually, the backend plugins could be responsible for cahching.
- Logic for inputs/outputs can be entirely decoupled from SG core, meaning
  there wouldn't even have to be new releases when additional inputs/outputs 
  are added
- Clear interface plugin API definition defines functionality of plugins and
  keeps feature creep from happening in core
- User Story:
    - I want to launch an SG for events with a loki backend. In startup.yaml (or
      passed in object of same format), I specify those ceilometer and loki
      plugins
- Developer's Story
    1) I need to make method for accepting and storing log messages into Loki
        - I build a loki plugin according to pre-defined interface
        - I build parsing and formating functionality in input plugin
        - Compile just my plugins, run specify them in the config and the SG
          core uses them
        - I can run some universal plugin tests on my plugins as well as
          implement my own
        - Can release new versions of plugins without waiting for a SG release
        - Can run new or updated plugin dynamically in already running Smart
          Gateway

    
NEXT STEPS (For monday when you get back):
1) Determine where caching will take place in new architecture
2) Draw diagrams for new architecture
3) Determine standardized plugin interface for inputs
4) Determine standardized plugin interface for database

