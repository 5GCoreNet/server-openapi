/*
 * nmbsf-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsf_MBSUserDataIngestSession

// Event - Possible values are: - USER_DATA_ING_SESS_STARTING: >     Indicates that the MBS User Data Ingest Session is starting. This is an \"MBS User Data      Ingest Session\" level event. - USER_DATA_ING_SESS_STARTED: >     Indicates that the MBS User Data Ingest Session started. This is an \"MBS User Data      Ingest Session\" level event. - USER_DATA_ING_SESS_TERMINATED: >     Indicates that the MBS User Data Ingest Session is terminated. This is an \"MBS User Data      Ingest Session\" level event. - DIST_SESS_STARTING: >     Indicates that the MBS Distribution Session is starting. This is an \"MBS Distribution      Session\" level event. - DIST_SESS_STARTED: >     Indicates that the MBS Distribution Session started. This is an \"MBS Distribution      Session\" level event. - DIST_SESS_TERMINATED: >     Indicates that the MBS Distribution Session is terminated. This is an \"MBS Distribution      Session\" level event. - DIST_SESS_SERV_MNGT_FAILURE: >     Indicates that the MBS Distribution Session could not be started (e.g. the necessary      resources could not be allocated by the MBS system). This is an \"MBS Distribution      Session\" level event. - DIST_SESS_POL_CRTL_FAILURE: >     Indicates that the MBS Distribution Session could not be started because of a policy      authorization/control failure or rejection. This is an \"MBS Distribution Session\"      level event. - DATA_INGEST_FAILURE: >     The MBS User Data Ingest is failed because the MBSTF is expecting data (the MBS Session      is active), but not receiving it. This is an \"MBS Distribution Session\" level event. - DELIVERY_STARTED: >     The MBS User Data delivery is started. - SESSION_TERMINATED: >     The MBS User Data Ingest Session is terminated. 
type Event struct {
}
