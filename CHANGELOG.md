# Change Log
All notable changes to this project will be documented in this file.

## [1.5.0] - 2017-03-22
### Added
- Encryption Everywhere support
- Certificates: Reissue and renewal of certificates
- Certificate Requests: Get and send notes linked to an order
- Certificate Requests: Send Subscriber Agreement
- Certificate Requests: You can now send a language parameter with the request, which will be used for (eg) the Subscriber Agreement
- Certificate Requests: A new field 'brandValidation' is added, which is used to indicate if the request is being held for review due to specific brand names

## [1.5.1] - 2023-11-22
### Added
- Country added to Request, Reissue, Renew structures
- Province added to Reques, Reissue, Renewt structures
- ApproverRepresentativeFirstName added to Request, Reissue, Renew structures
- ApproverRepresentativeLastName added to Request, Reissue, Renew structures
- ApproverRepresentativeEmail added to Request, Reissue, Renew structures
- ApproverRepresentativePhone added to Request, Reissue, Renew structures
- ApproverRepresentativePosition added to Request, Reissue, Renew structures

### Updated
- ScheduleValidationCall updated to accept various actions: ScheduledCallback, ManualCallback, ReplacePhone, replaceEmailAddress, sendCallbackEmail
