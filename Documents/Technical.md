# Technical Specifications

- Project: Smart Signage 6 App-solu

- Author: Paul MARIS

- Team:

| Role               | Name               |
|--------------------|--------------------|
| Project Manager    | Aurélien FERNANDEZ |
| Technical Leader   | Paul MARIS         |
| Program Manager    | Guillaume RIVIÈRE  |
| Software Engineer  | Lucien LAVATINE    |
| Q&A                | Lucas AUBARD       |

- Created on: 2022-09-29

- [Issue tracker](https://github.com/algosup/2022-2023-project-1-smart-signage-Project-6-group/issues)

## Table of content

<details>

  > **Note**
  > You can navigate through the document using the table of contents as shown below.
  > ![Tips](https://docs.github.com/assets/cb-47415/images/help/repository/headings_toc.png)

</details>

## 1. Introduction

### 1.1 Overview

Project ordered by [Signall](https://signall.com/), supervised by [ALGOSUP](https://algosup.com).

The goal of this project is to turn a sign on and off remotely. Being able to turn it on and off will also help with environmental issues. Signall also needs to know if the sign is on or off, and so they are able to know if the sign has issue(s). Currently, they have no other way than to go on site and check manually.

### 1.2 Problem description

Current problems encountered:

Signall:
- No way to know if the sign is on or off
- No way to turn the sign on or off remotely
- No way to know if the sign has issue(s)
- No way to automate the sign' lighting
- No way to automate/configure the sign' brightness

Final user:
- No other option to turn the sign on or off than unplugging the power cable
- Consumes more energy than necessary
- Takes time to turn the sign on or off

### 1.3 Glossary

| Word               | Definition                               |
|--------------------|------------------------------------------|
| Sign               | The LED Sign                             |
| LoraWAN            | The network protocol used to send and receive messages |
| ALGOSUP            | School where the project is made                  |

### 1.4 Context

The company Signall is looking for a prototype to make its electric signs intelligent, and control them remotely. They have contacted ALGOSUP to have a proof of concept/prototype.

### 1.5 Goals

This project aims to solve the following problems:
- Turn on and off a sign without having to plug out the power cable manually
    - The sign will be turned on and off automatically
- Detect, wether or not, issue(s) with the sign
    - And therefore, send an alert to the maintenance team, so they don't have to come regularly to check the proper functionning of the sign
- Know if the sign is off or not
- In an ecological approach, consume less electricity, optimizing the power consumption of the sign

### 1.6 Out of Scope

- Schedule the turn on and off of the sign
- Create a user interface to interact with the smart sign


### 1.7 Assumptions

Assumptions are future situations, beyond the control of the project, whose outcomes influence the success of a project. Constraints are limitations on the project that are beyond the control of the project team. The following assumptions and constraints apply to this project:

Assumptions:

- Availability of the hardware
- Developments in future technologies

Constraints:

- Governmental regulations
- Strategic decisions
- Budgetary constraints
- Competitive constraints

## 2. Solutions

### 2.1 Existing solutions

- [Signall](https://signall.com/), a company that sells electric signs. They have no way to know if the sign is on or off, and so they are not able to know if the sign has issue(s) yet, therefore they have no other way than to go on site and check manually.

- Timer Socket, a system that can turn on and off an electrical plug/outlet. It can be used to turn on and off a sign, but it is not a smart sign, and it is not connected anywhere.

- [Connected Sign](https://www.connectedsign.com/), a company that sells smart signs. They have a way to turn on and off the sign remotely, but they don't have a way to know if the sign has issue(s).

### 2.2 Proposed solution

##### The desired solution would work as follows:
- An electronic board *Lora e5 dev board*[^2], allowing to:
    - Receive *and* send data at long distance using radio waves at a low electrical cost. These data could contain:
        - The status of the sign (on or off)
        - Brightness of the sign (if relevant)
        - Issue status (if there is an issue or not)
        - Current timestamp (if relevant)
- Another electronic board, the *Blue Pill*[^3], allowing to:
    - Read the state of the sign (on or off) using a *relay*[^4].
    - Detect using *current sensor ZMCT103C*[^5] and the *current sensor ACS712*[^6] module if the sign has electrical irregularities/issue(s).

##### How it works together:

The *Blue Pill*[^3] connected physicaly to the *Lora e5 dev board*[^2], using a Serial communication the *Blue Pill*[^3] will send *AT commands* to the *Lora e5 dev board*[^2].
Those commands will be used to communicate with TheThingsNetwork[^7] thanks to the LoRaWAN protocol.

The solution will be powered on battery.

##### Data Model:

The amount of data that will be sent to the backend will be limited. Source: [TheThingsNetwork](https://www.thethingsnetwork.org/forum/t/limitation-of-transmission-of-data-in-one-day/36092/2)

> At the shortest distance from a Gateway you might be using a high data rate (SF7) and under the policy above your limit is approx 12,500 bytes of data per day.

> If your a long distnce from a Gateway you might be using a low data rate (SF12) and under the policy above your limit is approx 525 bytes of data per day.

The data will be sent from the LoRa to the Gateway ([TheThingsNetwork](https://console.cloud.thethings.network/)) and then will be sent to a server (out of scope) that will be able to store and send data.

- Receiving data
    The structure of the data has to be defined
- Send data
    The structure of the data has to be defined

##### Scalability:

The scalability of the project is not a problem, each individual electrical transformer (plugged to a sign) will need the solution. It means, one solution per electrical transformer as planned. According to the different characteristics of the signs, and in accordance with the customer's choice, a single sign may require several transformers depending on its size, consumption etc..
To conclude, a solution per electric transformer.

##### Fail case:

> Generally speaking a node does not know if a transmission has been successfully received. So it must act like it is with regards to fair access and duty cycles. For data redundancy it should act like the data might be lost and include some mechanism to recover from the loss (or ignore that possibility if a specific data point is not relevant and some loss is acceptable)

From [TheThingsNetwork](https://www.thethingsnetwork.org/forum/t/limitation-of-transmission-of-data-in-one-day/36092/2)

It means we have to design the system, keeping in mind that some data will be lost at some point.

### 2.3 Test plan

You can find the test plan [here](https://github.com/algosup/2022-2023-project-1-smart-signage-Project-6-group/blob/main/Documents/TestPlan.md), by Lucas AUBARD.

### 2.4 Alternative solutions

Few alternatives solutions have been considered, these solutions are similar in their overall functioning, however they include minor changes.

## 3. Structure considerations

### 3.1 Impact

The impact of the project is to make the electric signs of the company Signall more intelligent, and to be able to control them remotely. It means, a broader control of the signs, and a better management of the signs. For the maintenance team, it means less time spent on useless site (in case the sign has no issue).
For the customer, it means a better control of what's happening with the sign.

### 3.2 Cost analysis

Cost per day:
- Electric price
- Premium network

Product cost:
- 1 *Lora e5 dev board*[^2]: 27€
- 1 *Blue Pill*[^3]: 7€
- 1 *Relay*[^4]: 2€
- 1 *Current sensor ZMCT103C*[^5]: 3€
- 1 *Current sensor ACS712*[^6]: 3€
- 1 *Current transformer*: 4€
- 1 *Power cable*: 1€
- 1 *battery*: 3€

Energy cost:
- We need further information about the power consumption to predict with a higher accuracy the energy cost.

### 3.3 Security

### 3.4 Regional considerations

The solution is base on the LoRaWan protocol, which has a big coverage in France, using long range radio antennas.
However it is necessary to take into consideration the network coverage if the solution is exported to other countries.

### 3.5 Accessibility

Easy to install for the maintenance team.

### 3.6 Risks

- The network coverage is not good enough
- Rate limit of the network
- The sign is not plugged in the electrical transformer

## 4. Work

### 4.1 Timelines, work estimates, prioritization and milestones

The project timelines, work estimates, prioritization and milestones are available [here](https://github.com/orgs/algosup/projects/1).

## 5 References

[^1]: [Signall](https://signall.com/)

[^2]: LoRa E5 Wan Dev Kit [Official documentation and references](https://www.seeedstudio.com/LoRa-E5-Dev-Kit-p-4868.html)

[^3]: Blue Pill [Official documentation and references](https://stm32-base.org/boards/STM32F103C8T6-Blue-Pill.html)

[^4]: XY-MOS (relay) [Documentation and references](https://makershop.ie/xy-mos)

[^5]: ZMCT103C (precision current sensor) [Documentation and references](https://components101.com/sensors/zmct103c-precision-current-sensor-pinout-features-datasheet-alternative)

[^6]: ACS712 (current sensor) [Documentation and references](https://create.arduino.cc/projecthub/instrumentation-system/acs712-current-sensor-87b4a6)