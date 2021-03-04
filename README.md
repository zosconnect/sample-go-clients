# Go samples for demonstrating API orchestration using REST APIs on z/OS Connect
This repository contains Go samples that demonstrate combining data from multiple sources with API orchestration. The samples use REST APIs
created by [z/OS Connect Enterprise Edition](https://www.ibm.com/support/knowledgecenter/en/SS4SVW_3.0.0/com.ibm.zosconnect.doc/overview/what_is_new.html) to access z/OS applications and data hosted in subsystems such as CICS, IMS and Db2. Based off
of this [example](https://github.com/zosconnect/sample-nodejs-clients)

This code is licensed under an Apache 2.0 license. For details, see [license.txt](https://github.com/zosconnect/sample-go-clients/license.txt).

# Prerequisites
IBM Open Enterprise SDK for Go is installed and configured

# Installing
Clone this repository `git clone https://github.com/zosconnect/sample-go-clients`
To run any of the samples, `cd` to the appropriate directory and then type `go run <name_of_sample>`

Note that currently, each sample runs through a bunch of different inputs that are contained in `input.json` in each sample's folder. The file can be edited to test different cases.

# Sample 1: An Orchestration API that combines an IMS transaction with a Web API
This sample uses the [IMS phonebook application (IVTNO)](https://www.ibm.com/support/knowledgecenter/en/SSEPH2_15.1.0/com.ibm.ims15.doc.ins/ims_ivpsamples.htm). This phonebook application can *add a contact*, *delete a contact*,
*display a contact* and *update the contact*. The data is stored in an IMS database. In this sample, we will use the *display a contact*
function to list information about the contact based on the last name. The REST API for the transaction was created using [z/OS Connect
Enterprise Edition](https://www.ibm.com/support/knowledgecenter/en/SS4SVW_beta/com.ibm.zosconnect.doc/scenarios/ims_api_invoke.html)

Sample output
```
TOLOD
YVES
4166058936
L3R9Z7
43.818925
-79.330750
ON
Markham
CHATTERJEE
ABHISHEK
9054131234
L3R9Z7
43.818925
-79.330750
ON
Markham
MILLER
JAMES
9054131234
L3R9Z7
43.818925
-79.330750
ON
Markham
DOE
JOHN
9054131234
L3R9Z7
43.818925
-79.330750
ON
Markham
```

# Sample 2: A Microservice API that contains a health claim business rule
This sample is the Go microservice API that provides the health claim business rule that is used in the [z/OS Connect EE GitHub Sample on API requester](https://github.com/zosconnect/zosconnect-sample-cobol-apirequester). The sample provides automatic approval for a health claim based on the claim type and claim amount submitted. It handles the following claim types: **MEDICAL**, **DRUG**, **DENTAL**. The claim amount limits are **100 for MEDICAL**, **800 for DENTAL**, and **1000 for DRUG**. If the amount exceeded these limits, then the business rule will not approve the claim automatically.

Sample output
```
MEDICAL
100.000000
Accepted
Normal claim
MEDICAL
250.000000
Rejected
Amount exceeded $100. Claim requires further review
DENTAL
800.000000
Accepted
Normal claim
DENTAL
999.000000
Rejected
Amount exceeded $800. Claim requires further review
DRUG
1000.000000
Accepted
Normal claim
DRUG
3400.000000
Rejected
Amount exceeded $1000. Claim requires further review
```
