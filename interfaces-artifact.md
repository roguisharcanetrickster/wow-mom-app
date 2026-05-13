# "WOW Mom" Application Interfaces

This document defines the core API endpoints and Data Transfer Objects (DTOs) for the "WOW Mom - Mother Support Group Management Platform," aligning with the approved planning artifact.

## 1. Authentication Interfaces

### 1.1 User Authentication
- **`POST /auth/signup`**
  - Request: `{ email, password, role (Mother/Leader) }`
  - Response: `{ userId, token }`
- **`POST /auth/login`**
  - Request: `{ email, password }`
  - Response: `{ userId, token, role }`
- **`POST /auth/reset-password`**
  - Request: `{ email }` (initiates reset process)
  - Response: `{ message: "Password reset email sent" }`
- **`POST /auth/set-password`**
  - Request: `{ token, newPassword }` (completes reset process)
  - Response: `{ message: "Password updated successfully" }`

## 2. Mother Role Interfaces

### 2.1 Profile Management
- **`GET /mothers/me/profile`**
  - Response: `{ MotherProfileDto }`
- **`PUT /mothers/me/profile`**
  - Request: `{ FullName, PhoneNumber, Address, ZipCode, NumberOfChildren, ChildrenAges, PreferredMeetingTimes, PreferredLocations }`
  - Response: `{ MotherProfileDto }`

### 2.2 Group Discovery & Application
- **`GET /groups/available`**
  - Query Params: `location`, `meetingTime`, `capacityStatus`
  - Response: `[GroupDetailDto]`
- **`POST /groups/{groupId}/apply`**
  - Request: `{}`
  - Response: `{ applicationId, status: "Pending" }`
- **`GET /applications/me`**
  - Response: `[ApplicationSummaryDto]`
- **`GET /applications/{applicationId}`**
  - Response: `{ ApplicationDetailDto }`
- **`POST /applications/{applicationId}/confirm-interview`**
  - Request: `{ interviewTimeSlot }`
  - Response: `{ ApplicationDetailDto }`

## 3. Group Leader Role Interfaces

### 3.1 Group Management
- **`POST /leaders/me/groups`**
  - Request: `{ GroupName, GroupDescription, MeetingTime, MeetingLocation, MaxCapacity, MeetingFrequency }`
  - Response: `{ GroupDetailDto }`
- **`PUT /leaders/me/groups/{groupId}`**
  - Request: `{ GroupName, GroupDescription, MeetingTime, MeetingLocation, MaxCapacity, MeetingFrequency }`
  - Response: `{ GroupDetailDto }`
- **`GET /leaders/me/groups/{groupId}/roster`**
  - Response: `[MotherContactDto]` (for active members only)
- **`GET /leaders/me/groups`**
  - Response: `[GroupDetailDto]` (groups led by the leader)

### 3.2 Application Review & Interview
- **`GET /leaders/me/applications`**
  - Response: `[ApplicationSummaryDto]` (for applications to leader's groups)
- **`GET /leaders/me/applications/{applicationId}`**
  - Response: `{ ApplicationDetailDtoWithMotherProfile }`
- **`POST /leaders/me/applications/{applicationId}/schedule-interview`**
  - Request: `{ proposedTimeSlots: [DateTime] }`
  - Response: `{ ApplicationDetailDto }`
- **`POST /leaders/me/applications/{applicationId}/accept`**
  - Request: `{}`
  - Response: `{ ApplicationDetailDto }`
- **`POST /leaders/me/applications/{applicationId}/reject`**
  - Request: `{ rejectionReason (optional) }`
  - Response: `{ ApplicationDetailDto }`
- **`POST /leaders/me/applications/{applicationId}/activate`**
  - Request: `{}`
  - Response: `{ ApplicationDetailDto }`

### 3.3 Messaging (Optional)
- **`POST /leaders/me/mothers/{motherId}/message`**
  - Request: `{ messageContent }`
  - Response: `{ NotificationDto }`

## 4. Admin Role Interfaces

### 4.1 Dashboard & Overview
- **`GET /admin/dashboard/stats`**
  - Response: `{ totalMothers, totalGroups, acceptanceRate, capacityUtilization }`
- **`GET /admin/mothers`**
  - Query Params: `search`, `filter`
  - Response: `[MotherProfileDto]`
- **`GET /admin/groups`**
  - Query Params: `search`, `filter`
  - Response: `[GroupDetailDto]`
- **`GET /admin/applications`**
  - Query Params: `status`
  - Response: `[ApplicationDetailDto]`

### 4.2 Leader Management
- **`GET /admin/leaders/pending-approval`**
  - Response: `[LeaderProfileDto]`
- **`POST /admin/leaders/{leaderId}/approve`**
  - Request: `{}`
  - Response: `{ LeaderProfileDto }`
- **`POST /admin/leaders/{leaderId}/deactivate`**
  - Request: `{}`
  - Response: `{ LeaderProfileDto }`

### 4.3 System Settings & Reports
- **`POST /admin/system-settings`**
  - Request: `{ meetingTimeOptions, locationsList, groupSizeLimits }`
  - Response: `{ message: "Settings updated" }`
- **`GET /admin/reports/capacity`**
  - Response: `CSV file`
- **`GET /admin/reports/applications`**
  - Response: `CSV file`
- **`GET /admin/reports/demographics`**
  - Response: `CSV file`

## 5. Data Transfer Objects (DTOs) - Examples

### MotherProfileDto (based on MOTHERS table)
```json
{
  "motherId": "uuid",
  "email": "string",
  "fullName": "string",
  "phoneNumber": "string",
  "address": "string",
  "zipCode": "string",
  "numberOfChildren": "number",
  "childrenAges": "string[]", // or comma-separated string
  "preferredMeetingTimes": "string[]",
  "preferredLocations": "string[]",
  "registrationDate": "datetime",
  "accountStatus": "enum (Active, Inactive, Suspended)"
}
```

### GroupDetailDto (based on GROUPS table)
```json
{
  "groupId": "uuid",
  "groupName": "string",
  "groupDescription": "string",
  "leaderId": "uuid",
  "meetingTime": "enum (Morning, Afternoon, Evening)",
  "meetingLocation": "string",
  "maxCapacity": "number",
  "currentMemberCount": "number",
  "meetingFrequency": "enum (Weekly, Bi-weekly, Monthly)",
  "groupStatus": "enum (Active, Inactive, Full)"
}
```

### ApplicationDetailDto (based on APPLICATIONS table)
```json
{
  "applicationId": "uuid",
  "motherId": "uuid",
  "groupId": "uuid",
  "applicationDate": "datetime",
  "applicationStatus": "enum (Pending, Interview Scheduled, Accepted, Rejected, Active)",
  "interviewDate": "datetime | null",
  "interviewTimeSlot": "datetime | null",
  "interviewConfirmedByMother": "boolean",
  "rejectionReason": "string | null",
  "acceptedDate": "datetime | null",
  "activatedDate": "datetime | null",
  "notesFromLeader": "string | null"
}
```

### LeaderProfileDto (based on LEADERS table)
```json
{
  "leaderId": "uuid",
  "email": "string",
  "fullName": "string",
  "phoneNumber": "string",
  "leaderStatus": "enum (Active, Inactive, Pending Approval)",
  "approvedByAdmin": "boolean"
}
```

### NotificationDto (based on NOTIFICATIONS table)
```json
{
  "notificationId": "uuid",
  "userId": "uuid",
  "userType": "enum (Mother, Leader)",
  "notificationType": "string",
  "message": "string",
  "emailSent": "boolean",
  "read": "boolean",
  "createdAt": "datetime"
}
```

### MotherContactDto (for group roster)
```json
{
  "motherId": "uuid",
  "fullName": "string",
  "phoneNumber": "string",
  "email": "string"
}
```
