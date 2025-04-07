# Meetup in a Box - Build Invincible Apps with Durable Execution - Go

Welcome to the Temporal [Meetup in a Box](https://learn.temporal.io/meetup_in_a_box) program!
We're thrilled that you want to share Temporal within your community, and we hope these resources will aid you in doing just that.

This presentation is designed to introduce Go developers of all experience levels to Temporal.
In this presentation, you will lead the audience through the fundamental building blocks of Temporal, including the Temporal Service, Workflows, Activities, Workers, Durable Timers, and Task Queues.
You will explore a distributed Hello World example that uses an external service to obtain your IP address and then uses the result as input to a geo-location service, returning a greeting that provides your IP address and where you currently are.
Along the way, you'll demonstrate various failure scenarios, and how Temporal handles them by providing Durable Execution.

Thank you for your interest in Temporal, we hope you enjoy this Meetup in a Box!

## About this Repository

This repository contains the sample application for the _Build Invincible Apps with Durable Execution - Go_ presentation for Meetup in a box.
For more information about this presentation, visit the presentations [landing page](https://learn.temporal.io/meetup_in_a_box/invincible_apps/go/).

## Installation

In order to present this material, you will need to install a development version of the Temporal Service, the Temporal Go SDK, and various packages onto your presentation device.

1. Install the Temporal CLI on your local machine by following [this tutorial](https://learn.temporal.io/getting_started/go/dev_environment/#set-up-a-local-temporal-service-for-development-with-temporal-cli)
2. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/temporal-community/miab-build-invincible-apps-go.git
    ```
3. Change directories into the cloned repository:
    ```bash
    cd miab-build-invincible-apps-go
    ```
4. Change directories into the `iplocate` directory
    ```bash
    cd iplocate
    ```
5. Install the Go dependencies
    ```bash
    `go mod tidy`
    ```

## Teaser Demo

This demo is a quick, 1 - 2 minute teaser of Durable Execution. 
It should _not_ be used to introduce any Temporal Concepts yet, that will come later in the presentation. 

This demo will quickly demonstrate the difference between "" and durable execution.
Normal execution is any execution that is not durable, so it is synonymous with the majority of code that is executed by most developers on a daily basis. 

This demo will run a loop that counts from 1 to 10, sleeping for 1 second between each iteration, and logging the value to the screen.
After the loop progresses a bit (about 4 or 5), the loop will be terminated.
In a normal execution, restarting this process will cause the loop to start over from the beginning, but in a durable process we'll see that the execution resumes where it was, continuing the execution.

### Running the demo

The sample application requires three terminal windows and a browser to run. 
You can either open three separate terminals, or use a terminal multiplex such as `screen` or `tmux` to manage your terminals.

1. In the first terminal, run the following command to start the Temporal Service on port 8080 with a persistent database:
    ```bash
    temporal server start-dev --ui-port 8080 --db-filename temporal.db
    ```
    You should see an output _similar_ to this (the exact output and versions may differ)
    ```bash
    CLI 1.1.1 (Server 1.25.1, UI 2.31.2)

    Server:  localhost:7233
    UI:      http://localhost:8080
    Metrics: http://localhost:65134/metrics
    ```
2. In the second terminal, run the following commands to demo Normal Execution.
The Temporal Service you started will not be used in this portion.
    1. Change directories into the `durable-vs-normal-execution/normal` directory:
        ```bash
        cd durable-vs-normal-execution/normal
        ```
    2. Show the code to the audience, demonstrating it's a standard loop. Then run the Normal Execution script
        ```bash
        go run main.go
        ```
    3. The script should log the numbers 1 to 10 one at a time, sleeping for 1 second in between, similar to this:
        ```
        0
        1
        2
        3
        4
        ```
        Once it has gotten to 4 or 5, use `CTRL-C` and kill the execution.
   
    4. Ask the audience 
        > What will happen if I run the execution again?
    
        The answer should be "It will start over."
        You can ask

        > Who thinks it will resume where it left off?"

        No one will probably think this, as this is not typical behavior for any programming language.
   
    5. Run the Normal Execution script again
        ```bash
        go run main.go
        ```
        Confirm that the code did, in fact, start over.
   
4. For this demo you will need both the second and third terminal. Now you will demonstrate the same application, a loop that counts from 1 to 10, sleeping between each step, that does survive killing the process. When the execution is brought back online, it will resume where it left off. 

   **Remember**, this is not the time to go deep into the architecture of Temporal.

   Mention that for this sample, there's a running process that executes the code and another script that invokes it. 
    1. Change directories into the `durable-vs-normal-execution/durable` directory:
        ```bash
        cd durable-vs-normal-execution/durable
        ```
    2. Show the `workflow.go` and `activities.go` file to the audience. 
        Explain that it's a slight modification, but relatively similar.
        You added a function to do the addition. 
        **DON'T** go into depth here, all the terms will be defined later.
    3. Start the Temporal Worker:
        ```bash
        go run worker/main.go
        ```

        You should see output _similar_ to this:
        ```bash
        Started Worker
        ```
    4. Run the starter script, which will begin the execution of the loop code
        ```bash
        go run starter/main.go
        ```
        This code will exit immediately. Return to the other terminal.
    6. In the previous terminal you should see the output of the loop. Once the loop gets to 4 or 5, CTRL-C the process and kill it. You will see extra output:
        ```
        ^CWorker has been stopped.
        ```
        For now you can ignore these.
    7. Ask the same question from the previous demo:

        > What will happen if I run the execution again?
    
        See if the audience answers differently. But then add:
    
        > Wouldn't it be great if it did just resume working?

    8. Restart the Worker and let the code complete
        ```bash
        go run worker/main.go
        ```
        The code should pick up where it left off, resuming from the next number in the count.
   
    9. Open a browser tab to [http://127.0.0.1:8080](http://127.0.0.1:8080) to view the Temporal Web UI. Go to the Workflow Execution of this run and show the timeline view. You should see the timers and the call to the `addOne` function. Click on a single Activity Task and show the input and output. 

    People may be very interested and start asking a lot of questions, try to hold them off, there will be plenty of time going forward. Go back to the slides and resume the presentation.

    
## Temporal Demos

Below are the instructions to the demos to showcase various features of Temporal using the distributed Hello World application.
Before doing these demos, ensure that you are connected to a network.
These demos are listed in a chronological order, meaning they are designed to be done one after the other, each one building upon the previous demo.
While presenting, feel free to elaborate within these demos.
Personal anecdotes and experiences make the presentation more relatable to the audience.
You will be entering the name of a person into the form multiple times, so to encourage participation, use audience member's' names in the form as input.
While you are demoing, narrate what you are doing as you are doing it.
The more you speak, the more engaged the audience will be.
We have provided scripts in areas below if you need them, but feel free to explain it in your own way if you feel comfortable doing so.

### Running the Application

The sample application requires three terminal windows and a browser to run. 
You can either open three separate terminals, or use a terminal multiplex such as `screen` or `tmux` to manage your terminals.

1. In the first terminal, run the following command to start the Temporal Service on port 8080 with a persistent database:
    ```bash
    temporal server start-dev --ui-port 8080 --db-filename temporal.db
    ```
    You should see an output _similar_ to this (the exact output and versions may differ)
    ```
    CLI 1.1.1 (Server 1.25.1, UI 2.31.2)

    Server:  localhost:7233
    UI:      http://localhost:8080
    Metrics: http://localhost:65134/metrics
    ```
2. In the second terminal, run the following commands to start the Temporal Worker:
    1. Change directories into the `iplocate` directory:
        ```bash
        cd iplocate
        ```
    2. Start the Temporal Worker:
        ```bash
        go run worker/main.go
        ```
3. In the third terminal, run the following commands to start the web application:
    1. Change directories into the `iplocate` directory:
        ```bash
        cd iplocate
        ```
    2. Start the web application on port `8000`:
        ```bash
        go run server/main.go
        ```
4. Open a browser tab to [http://127.0.0.1:8000](http://127.0.0.1:8000) to view the web application.
5. Open a browser tab to [http://127.0.0.1:8080](http://127.0.0.1:8080) to view the Temporal Web UI.

### Demo #1 Successful Execution

**Purpose**: Show the working application  
**Failure**: None  
**Temporal Feature Demonstrated**: Temporal Web UI  
**Audience Takeaway**: Temporal keeps track of the results of each Activity and Workflow.
We can use the Web UI to debug issues when they go wrong.

1. Prepare the audience with an introduction to the demo:
    > "First, let's see the application work successfully. This application is a distributed take on a Hello World program. It takes a user's name, and makes HTTP calls to two different services. The first service retrieves the IP address of the requester and returns it, and the second service takes an IP address and uses it to perform a geo-lookup. This means that the result of the first request will be used as the input for the second request. Once both requests have completed, the application will return a greeting, saying hello, your IP address, and what city you are currently in."
2. Enter your name in the text field **Enter your name** in the web application and press the **Get Greeting** button.
3. The app will return a greeting immediately similar to the following:
    ```
    Hello, Ziggy!
    Your IP Address is 256.256.256.256.
    You are in Austin, Texas, United States
    ```
4. Open the Web UI and walk through the Event History, highlighting the following things:
    1. The input and output to the Workflow
    2. The input and output to each Activity, noting how the input to the `getLocationInfo` Activity was the output from the `getIP` Activity.
    3. The timeline view
    4. Anything else you find interesting that you wish to point out, or any anecdotes on how you use the Web UI
5. Return to the web application and click the **Reset** button.

### Demo #2 Successful Execution with a Durable Timer

**Purpose**: Show the working application with a small pause happening between the service calls  
**Failure**: None  
**Temporal Feature Demonstrated**: Durable Timers  
**Audience Takeaway**: Temporal supports Timers. This is setting up for the next demo and allowing us to create time to simulate failures.  

1. Explain to the audience that next you'll introduce a delay in between the two calls:
    > "Next, I'm going to introduce a delay between the two calls. This is going to help me simulate various failure scenarios going forward, and also demonstrate the concept of Durable Timers. So this application can take in an integer via the form and sleep for that amount of seconds after executing the first Activity but before executing the second."
2. Enter an audience member's name in the text field **Enter your name** in the web application.
3. Press the **Show Demo Options** link on the page.
4. In the **Sleep Duration (seconds)** section of the form, provide **3** in the **Number of Seconds** field and press the **Get Greeting** button.
5. The greeting should return after a short delay (3 seconds) similar to the original:
    ```
    Hello, Ziggy!
    Your IP Address is 256.256.256.256.
    You are in Austin, Texas, United States
    ```
6. Open the Web UI and investigate in the Event History. Locate the **Timer Started** and **Timer Fired** events
    1. Tell the audience that Timers are managed on the Temporal Service side, and will fire regardless if there is a Worker running or not.
    If a Timer fires and a Worker is not available, it will pick up when a Worker becomes available again.
7. Return to the web application and click the **Reset** button.

### Demo #3 Worker Becomes Unavailable During an Execution

**Purpose**: Demonstrate that a Timer fires regardless if there is a Worker running.  
**Failure**: Worker will "become unavailable" (ctrl-c) while a Timer is waiting to fire  
**Temporal Feature Demonstrated**: Durable Timers  
**Audience Takeaway**: Timers are managed on the Temporal Service side. 
They will fire regardless if there is a Worker running or not.
If a Timer fires and a Worker is not available, it will pick up when a Worker becomes available again.

1. Explain to the audience that a Durable Timer will fire regardless if a Worker is available:
    > "But what happens if the Worker were to become unavailable before the Timer were to fire? If this wasn't a Durable Timer, the state of the Timer would be lost, and you would have to restart the application and wait for the entire timer again. With Temporal, the Timer will fire regardless if there is a Worker running, and it will resume execution as soon as a new Worker comes online."
2. Enter an audience member's name in the text field **Enter your name** in the web application.
3. Press the **Show Demo Options** link on the page.
4. In the **Sleep Duration (seconds)** section of the form, provide **10** in the **Number of Seconds** field and press the **Get Greeting** button.
5. Immediately after, switch to the terminal with the Worker running, and press `CTRL-C` (or `CMD-C` if on Mac) to kill the Worker process.
    1. You want to be sure to kill the Worker before the Timer fires. 
6. Wait ~10 seconds and show the audience that the results haven't appeared on the screen.
7. Go to the Web UI and view the Event History.
    1. Show the audience that the **Timer Started** and **Timer Fired** events have both been recorded, but nothing else has.
    2. Explain to the audience this is because there is no Worker running.
8. Go back to the terminal and restart the Worker:
    ```bash
    go run worker/main.go
    ```
9. The greeting should immediately return similar to the original:
    ```
    Hello, Ziggy!
    Your IP Address is 256.256.256.256.
    You are in Austin, Texas, United States
    ```
10. Open the Web UI and show that the Workflow continued execution and everything looks normal, as if nothing ever happened.
11. Explain to the audience:
    > "The Worker resumed execution as if nothing happened. The first Activity was not re-executed. The state of the application was reconstructed from the Event History, and the result that was returned from the successful execution of the Activity the first time was used. The first Activity was not re-executed."
12. Ask the audience the following question:
    > "Now lets say the Timer had been set to longer, for example, an hour, and we had recovered the Worker before it had fired, what would have happend?"
    **Answer:** The Timer would have fired and execution would have completed successfully, as if nothing had ever happened.
13. Return to the web application and click the **Reset** button.

### Demo #4 Temporal Outage During an Execution

**Purpose**: Demonstrate that the Temporal Service going down won't stop us!  
**Failure**: The Temporal Service will "crash" (ctrl-c) while the Timer is waiting to fire.  
**Temporal Feature Demonstrated**: Durability of Temporal  
**Audience Takeaway**: The Temporal Service itself is durable.  

1. Explain to the audience that a Durable Timer will fire regardless if a Worker is available:
    > "You've seen the Worker become unavailable, you've seen the network fail. But what if Temporal fails? "
2. Enter an audience member's name in the text field **Enter your name** in the web application.
3. Press the **Show Demo Options** link on the page.
4. In the **Sleep Duration (seconds)** section of the form, provide **10** in the **Number of Seconds** field and press the **Get Greeting** button.
5. Go to the Temporal Web UI home page and see that there is a Workflow in the **Running** state.
6. Immediately after, switch to the terminal with the Temporal Service running, and press `CTRL-C` (or `CMD-C` if on Mac) to kill the Temporal Service.
    1. You want to be sure to kill the Service before the Timer fires. 
7. Wait ~10 seconds and show the audience that the results haven't appeared on the screen.
8. Go to the Web UI and show that it is refusing to connect.
9. Go back to the terminal and restart the Temporal Service:
    ```bash
    temporal server start-dev --ui-port 8080 --db-filename temporal.db
    ```
10. Refresh the Temporal WebUI and observe that the Workflow has completed.
11. Return to the application and the greeting should have appeared as well.
12. Open the Web UI and show that the Workflow continued execution and everything looks normal, as if nothing ever happened.
13. Conclude your demo to the audience:
    > "As long as the database hasn't been corrupted or deleted, the Temporal Service resumes right where it left off. You may have correctly deduced then that managing that persistence layer is one of the most important parts of running a self-hosted Temporal instance. But no matter what we throw at it, Temporal survived it."
Comment

### Demo #5 - Recovering from a Bug

**Purpose**: Demonstrate that Temporal allows you to fix bugs even while the code is running  
**Failure**: We're going to introduce a bug into our code, fix it, and watch Temporal recover  
**Temporal Feature Demonstrated**: Durable execution  
**Audience Takeaway**: Temporal allows developers to address bugs while the code is running.  

1. Explain to the audience that Temporal can even let you fix bugs while the code is running, without having to stop execution:
    > "If you have a bug in your code, what do you usually have to do? Stop the code, fix it, and then restart the code. But what about in Temporal? If Activities are retried forever, and the previous state is saved, shouldn't we be able to fix the code, deploy it, and continue execution?"
2. Open a text editor and modify `activities.go`:
    1. On line 39, modify the URL by removing the `h` in `http`. This will create an invalid URL, causing the Activity to fail.
    2. Restart your Worker:
        ```bash
        go run worker/main.go
        ```
3. Enter an audience member's name in the text field **Enter your name** in the web application.
4. Observe that a result was not displayed.
5. Go to the Temporal Web UI and you'll see the `getLocationInfo` Activity is failing due to the invalid URL supplied.
6. Return to your text editor and modify `activities.go`, adding the `h` back to `http`.
7. Restart the Worker by first killing it with `CTRL-C` and then running:
    ```bash
       go run worker/main.go
    ```
8. Wait until the **Next Retry** time passes, in which the change should be picked up, and the Activity should complete successfully.
9. Switch back to the web application, where you should see the response.
10. Open the Web UI and show that the Workflow continued execution and everything looks normal, as if nothing ever happened.
11. Conclude your demo to the audience:
    > "If you have a bug in an Activity, and that Activity is failing, you can fix the bug and redeploy it, and Temporal will pick up this change and continue executing without losing the progress that was made previously."


**You have now completed the demo**

### Bonus Demo - Network Outage During an Execution
**Warning**: If you are giving a virtual/hybrid meetup, this demo is not possible. 
If the venue uses a captive portal WiFi, reconnecting may be difficult. 
Only do this demo if you can disconnect without causing an interruption, and trust that it will reconnect without issue. 

**Purpose**: Demonstrate how Temporal retries until either success or cancellation  
**Failure**: Network will have an "outage" as the presenter will turn off their Wifi while the app is running.  
**Temporal Feature Demonstrated**: Activity Retries  
**Audience Takeaway**: Temporal automatically retries Activities on failure.  
An intermittent failure, such as a network outage, will be retried until the Activity completes successfully or is cancelled.

1. Ask the audience what they think will happen if the network were to go out while the application was running:
    > "We all know that networks can be unreliable. So what do you think would happen if during the delay, I simulated a network failure by turning off my Wifi?"
2. Enter an audience member's name in the text field **Enter your name** in the web application.
3. Press the **Show Demo Options** link on the page.
4. In the **Sleep Duration (seconds)** section of the form, provide **10** in the **Number of Seconds** field and press the **Get Greeting** button.
5. Immediately after, disable the network on your device, either by turning off the Wifi or unplugging the ethernet cable.
6. Wait for the Timer to fire and the second Activity to begin attempting to execute.
    1. You will notice that the Activity does not complete, as the network is down and the service cannot be reached.
7. Go to the Web UI and locate the **Pending Activities** line in the Event History. Click on it to expand it.
    1. Show the audience the **Attempt** count, the time until **Next Retry**, and the **Last Failure** message.
8. Restore the network on your device. 
9. After the **Next Retry** time has passed, the greeting should return similar to the original:
    ```
    Hello, Ziggy!
    Your IP Address is 256.256.256.256.
    You are in Austin, Texas, United States
    ```
10. Explain to the audience:
    > "Activities are retried by default in Temporal. How often, how long, how many times can be configured, but intermittent errors don't cause crashes."
12. Return to the web application and click the **Reset** button.

