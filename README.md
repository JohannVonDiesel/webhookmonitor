## What is this mess?

WIP!  demo webhook monitor that generates an event when a POST request is received, it generates a response and also loggs the events for further data injestion.  
It is intended to be used as a basis for building a more robust monitoring scheme for applications or services that generate webhook events.  
  
This deploys as a GCP Cloud Function.  
  
There is a deployment script that can be modified.  It might be a good idea to set up a go.env and call it for certain vars.  