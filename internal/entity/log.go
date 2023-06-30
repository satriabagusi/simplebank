/*
Author: Satria Bagus(satria.bagus18@gmail.com)
log.go (c) 2023
Desc: description
Created:  2023-06-30T11:31:44.097Z
Modified: !date!
*/

package entity

import "time"

type Log struct {
	ID         string    `json:"id"`
	LogName    string    `json:"log_name"`
	LogStatus  string    `json:"log_status"`
	LogMessage string    `json:"log_message"`
	Timestamp  time.Time `json:"timestamp"`
}
