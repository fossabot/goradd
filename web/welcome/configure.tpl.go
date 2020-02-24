//** This file was code generated by got. DO NOT EDIT. ***

package welcome

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/pkg/resource"
)

func init() {
	resource.RegisterPath("/goradd/configure.html",
		func(ctx context.Context, buf *bytes.Buffer) (headers map[string]string, err error) {

			buf.WriteString(`<!DOCTYPE html>
<head>
<meta charset="utf-8"/>
<title>`)

			buf.WriteString(`Configuring the Database`)

			buf.WriteString(`</title>
</head>
<body>
`)

			buf.WriteString(`
<h1>Configuring the Database</h1>
<p>
	Goradd relies on your database(s) not only to get to data, but to understand its structure. For
	sql databases, it does this by reading the database schema. When NoSQL databases are supported,
	the structure will come from a configuration file.
</p>
<p>
	Using its knowledge of the structure of your database, the goradd code generator will
	create data models, data nodes, forms, and more to get you a working application to start customizing.
</p>

<p>
	Goradd currently only supports MySQL databases. Adapters to other databases are easy to write, so
	if you would like to see another database supported and are willing to help, open an issue
	at our github site.
</p>
<p>
	To configure your database, edit the goradd-project/config/db.go file. You will see directions there.
	If you want to run the examples code and tutorial, you should create a local copy of the goradd
	examples database. To do that, create a "goradd" database and import the sql code found in the goradd/web/examples/db directory,
	which is reprinted below for your convenience.
</p>
<code>
`)

			buf.WriteString(`<p>-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Jan 27, 2020 at 01:59 AM
-- Server version: 5.7.28
-- PHP Version: 7.2.19</p>
<p>SET FOREIGN_KEY_CHECKS=0;
SET SQL_MODE = &#34;NO_AUTO_VALUE_ON_ZERO&#34;;
SET time_zone = &#34;+00:00&#34;;</p>
<p>--
-- Database: ` + "`" + `goradd` + "`" + `
--</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `address` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `address` + "`" + ` (
                           ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                           ` + "`" + `person_id` + "`" + ` int(11) UNSIGNED NOT NULL,
                           ` + "`" + `street` + "`" + ` varchar(100) NOT NULL,
                           ` + "`" + `city` + "`" + ` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `address` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `address` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `person_id` + "`" + `, ` + "`" + `street` + "`" + `, ` + "`" + `city` + "`" + `) VALUES
(1, 1, &#39;1 Love Drive&#39;, &#39;Phoenix&#39;),
(2, 2, &#39;2 Doves and a Pine Cone Dr.&#39;, &#39;Dallas&#39;),
(3, 3, &#39;3 Gold Fish Pl.&#39;, &#39;New York&#39;),
(4, 3, &#39;323 W QCubed&#39;, &#39;New York&#39;),
(5, 5, &#39;22 Elm St&#39;, &#39;Palo Alto&#39;),
(6, 7, &#39;1 Pine St&#39;, &#39;San Jose&#39;),
(7, 7, &#39;421 Central Expw&#39;, &#39;Mountain View&#39;);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `employee_info` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `employee_info` + "`" + ` (
                                 ` + "`" + `id` + "`" + ` int(11) NOT NULL,
                                 ` + "`" + `person_id` + "`" + ` int(11) UNSIGNED NOT NULL,
                                 ` + "`" + `employee_number` + "`" + ` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `login` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `login` + "`" + ` (
                         ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                         ` + "`" + `person_id` + "`" + ` int(11) UNSIGNED DEFAULT NULL,
                         ` + "`" + `username` + "`" + ` varchar(20) NOT NULL,
                         ` + "`" + `password` + "`" + ` varchar(20) DEFAULT NULL,
                         ` + "`" + `is_enabled` + "`" + ` tinyint(1) NOT NULL DEFAULT &#39;1&#39;
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `login` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `login` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `person_id` + "`" + `, ` + "`" + `username` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `is_enabled` + "`" + `) VALUES
(1, 1, &#39;jdoe&#39;, &#39;p@$$.w0rd&#39;, 0),
(2, 3, &#39;brobinson&#39;, &#39;p@$$.w0rd&#39;, 1),
(3, 4, &#39;mho&#39;, &#39;p@$$.w0rd&#39;, 1),
(4, 7, &#39;kwolfe&#39;, &#39;p@$$.w0rd&#39;, 0),
(5, NULL, &#39;system&#39;, &#39;p@$$.w0rd&#39;, 1);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `milestone` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `milestone` + "`" + ` (
                             ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL,
                             ` + "`" + `project_id` + "`" + ` int(10) UNSIGNED NOT NULL,
                             ` + "`" + `name` + "`" + ` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `milestone` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `milestone` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `project_id` + "`" + `, ` + "`" + `name` + "`" + `) VALUES
(1, 1, &#39;Milestone A&#39;),
(2, 1, &#39;Milestone B&#39;),
(3, 1, &#39;Milestone C&#39;),
(4, 2, &#39;Milestone D&#39;),
(5, 2, &#39;Milestone E&#39;),
(6, 3, &#39;Milestone F&#39;),
(7, 4, &#39;Milestone G&#39;),
(8, 4, &#39;Milestone H&#39;),
(9, 4, &#39;Milestone I&#39;),
(10, 4, &#39;Milestone J&#39;);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `person` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `person` + "`" + ` (
                          ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                          ` + "`" + `first_name` + "`" + ` varchar(50) NOT NULL,
                          ` + "`" + `last_name` + "`" + ` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `person` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + `) VALUES
(1, &#39;John&#39;, &#39;Doe&#39;),
(2, &#39;Kendall&#39;, &#39;Public&#39;),
(3, &#39;Ben&#39;, &#39;Robinson&#39;),
(4, &#39;Mike&#39;, &#39;Ho&#39;),
(5, &#39;Alex&#39;, &#39;Smith&#39;),
(6, &#39;Wendy&#39;, &#39;Smith&#39;),
(7, &#39;Karen&#39;, &#39;Wolfe&#39;),
(8, &#39;Samantha&#39;, &#39;Jones&#39;),
(9, &#39;Linda&#39;, &#39;Brady&#39;),
(10, &#39;Jennifer&#39;, &#39;Smith&#39;),
(11, &#39;Brett&#39;, &#39;Carlisle&#39;),
(12, &#39;Jacob&#39;, &#39;Pratt&#39;);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `person_persontype_assn` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `person_persontype_assn` + "`" + ` (
                                          ` + "`" + `person_id` + "`" + ` int(11) UNSIGNED NOT NULL,
                                          ` + "`" + `person_type_id` + "`" + ` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `person_persontype_assn` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `person_persontype_assn` + "`" + ` (` + "`" + `person_id` + "`" + `, ` + "`" + `person_type_id` + "`" + `) VALUES
(3, 1),
(10, 1),
(1, 2),
(3, 2),
(7, 2),
(1, 3),
(3, 3),
(9, 3),
(2, 4),
(7, 4),
(2, 5),
(5, 5);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `person_type` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `person_type` + "`" + ` (
                               ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                               ` + "`" + `name` + "`" + ` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `person_type` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `person_type` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `) VALUES
(4, &#39;Company Car&#39;),
(1, &#39;Contractor&#39;),
(3, &#39;Inactive&#39;),
(2, &#39;Manager&#39;),
(5, &#39;Works From Home&#39;);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `person_with_lock` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `person_with_lock` + "`" + ` (
                                    ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                                    ` + "`" + `first_name` + "`" + ` varchar(50) NOT NULL,
                                    ` + "`" + `last_name` + "`" + ` varchar(50) NOT NULL,
                                    ` + "`" + `sys_timestamp` + "`" + ` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `person_with_lock` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `person_with_lock` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + `, ` + "`" + `sys_timestamp` + "`" + `) VALUES
(1, &#39;John&#39;, &#39;Doe&#39;, NULL),
(2, &#39;Kendall&#39;, &#39;Public&#39;, NULL),
(3, &#39;Ben&#39;, &#39;Robinson&#39;, NULL),
(4, &#39;Mike&#39;, &#39;Ho&#39;, NULL),
(5, &#39;Alfred&#39;, &#39;Newman&#39;, NULL),
(6, &#39;Wendy&#39;, &#39;Johnson&#39;, NULL),
(7, &#39;Karen&#39;, &#39;Wolfe&#39;, NULL),
(8, &#39;Samantha&#39;, &#39;Jones&#39;, NULL),
(9, &#39;Linda&#39;, &#39;Brady&#39;, NULL),
(10, &#39;Jennifer&#39;, &#39;Smith&#39;, NULL),
(11, &#39;Brett&#39;, &#39;Carlisle&#39;, NULL),
(12, &#39;Jacob&#39;, &#39;Pratt&#39;, NULL);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `project` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `project` + "`" + ` (
                           ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                           ` + "`" + `num` + "`" + ` int(11) NOT NULL COMMENT &#39;To simplify checking test results and as a non pk id test&#39;,
                           ` + "`" + `project_status_type_id` + "`" + ` int(11) UNSIGNED NOT NULL,
                           ` + "`" + `manager_id` + "`" + ` int(11) UNSIGNED DEFAULT NULL,
                           ` + "`" + `name` + "`" + ` varchar(100) NOT NULL,
                           ` + "`" + `description` + "`" + ` text,
                           ` + "`" + `start_date` + "`" + ` date DEFAULT NULL,
                           ` + "`" + `end_date` + "`" + ` date DEFAULT NULL,
                           ` + "`" + `budget` + "`" + ` decimal(12,2) DEFAULT NULL,
                           ` + "`" + `spent` + "`" + ` decimal(12,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `project` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `num` + "`" + `, ` + "`" + `project_status_type_id` + "`" + `, ` + "`" + `manager_id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `start_date` + "`" + `, ` + "`" + `end_date` + "`" + `, ` + "`" + `budget` + "`" + `, ` + "`" + `spent` + "`" + `) VALUES
(1, 1, 3, 7, &#39;ACME Website Redesign&#39;, &#39;The redesign of the main website for ACME Incorporated&#39;, &#39;2004-03-01&#39;, &#39;2004-07-01&#39;, &#39;9560.25&#39;, &#39;10250.75&#39;),
(2, 2, 1, 4, &#39;State College HR System&#39;, &#39;Implementation of a back-office Human Resources system for State College&#39;, &#39;2006-02-15&#39;, NULL, &#39;80500.00&#39;, &#39;73200.00&#39;),
(3, 3, 1, 1, &#39;Blueman Industrial Site Architecture&#39;, &#39;Main website architecture for the Blueman Industrial Group&#39;, &#39;2006-03-01&#39;, &#39;2006-04-15&#39;, &#39;2500.00&#39;, &#39;4200.50&#39;),
(4, 4, 2, 7, &#39;ACME Payment System&#39;, &#39;Accounts Payable payment system for ACME Incorporated&#39;, &#39;2005-08-15&#39;, &#39;2005-10-20&#39;, &#39;5124.67&#39;, &#39;5175.30&#39;);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `project_status_type` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `project_status_type` + "`" + ` (
                                       ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL,
                                       ` + "`" + `name` + "`" + ` varchar(50) NOT NULL,
                                       ` + "`" + `description` + "`" + ` text,
                                       ` + "`" + `guidelines` + "`" + ` text,
                                       ` + "`" + `is_active` + "`" + ` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `project_status_type` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `project_status_type` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `guidelines` + "`" + `, ` + "`" + `is_active` + "`" + `) VALUES
(1, &#39;Open&#39;, &#39;The project is currently active&#39;, &#39;All projects that we are working on should be in this state&#39;, 1),
(2, &#39;Cancelled&#39;, &#39;The project has been canned&#39;, NULL, 1),
(3, &#39;Completed&#39;, &#39;The project has been completed successfully&#39;, &#39;Celebrate successes!&#39;, 1),
(4, &#39;Planned&#39;, &#39;Project is in the planning stages and has not been assigned a manager&#39;, &#39;Get ready&#39;, 0);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `related_project_assn` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `related_project_assn` + "`" + ` (
                                        ` + "`" + `parent_id` + "`" + ` int(11) UNSIGNED NOT NULL,
                                        ` + "`" + `child_id` + "`" + ` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `related_project_assn` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `related_project_assn` + "`" + ` (` + "`" + `parent_id` + "`" + `, ` + "`" + `child_id` + "`" + `) VALUES
(4, 1),
(1, 3),
(1, 4);</p>
<p>-- --------------------------------------------------------</p>
<p>--
-- Table structure for table ` + "`" + `team_member_project_assn` + "`" + `
--</p>
<p>CREATE TABLE ` + "`" + `team_member_project_assn` + "`" + ` (
                                            ` + "`" + `team_member_id` + "`" + ` int(11) UNSIGNED NOT NULL,
                                            ` + "`" + `project_id` + "`" + ` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;</p>
<p>--
-- Dumping data for table ` + "`" + `team_member_project_assn` + "`" + `
--</p>
<p>INSERT INTO ` + "`" + `team_member_project_assn` + "`" + ` (` + "`" + `team_member_id` + "`" + `, ` + "`" + `project_id` + "`" + `) VALUES
(2, 1),
(5, 1),
(6, 1),
(7, 1),
(8, 1),
(2, 2),
(4, 2),
(5, 2),
(7, 2),
(9, 2),
(10, 2),
(1, 3),
(4, 3),
(6, 3),
(8, 3),
(10, 3),
(1, 4),
(2, 4),
(3, 4),
(5, 4),
(8, 4),
(11, 4),
(12, 4);</p>
<p>--
-- Indexes for dumped tables
--</p>
<p>--
-- Indexes for table ` + "`" + `address` + "`" + `
--
ALTER TABLE ` + "`" + `address` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD KEY ` + "`" + `IDX_address_1` + "`" + ` (` + "`" + `person_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `employee_info` + "`" + `
--
ALTER TABLE ` + "`" + `employee_info` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD UNIQUE KEY ` + "`" + `person_id` + "`" + ` (` + "`" + `person_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `login` + "`" + `
--
ALTER TABLE ` + "`" + `login` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD UNIQUE KEY ` + "`" + `IDX_login_2` + "`" + ` (` + "`" + `username` + "`" + `),
    ADD UNIQUE KEY ` + "`" + `IDX_login_1` + "`" + ` (` + "`" + `person_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `milestone` + "`" + `
--
ALTER TABLE ` + "`" + `milestone` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD KEY ` + "`" + `IDX_milestoneproj_1` + "`" + ` (` + "`" + `project_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `person` + "`" + `
--
ALTER TABLE ` + "`" + `person` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD KEY ` + "`" + `IDX_person_1` + "`" + ` (` + "`" + `last_name` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `person_persontype_assn` + "`" + `
--
ALTER TABLE ` + "`" + `person_persontype_assn` + "`" + `
    ADD PRIMARY KEY (` + "`" + `person_id` + "`" + `,` + "`" + `person_type_id` + "`" + `),
    ADD KEY ` + "`" + `person_type_id` + "`" + ` (` + "`" + `person_type_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `person_type` + "`" + `
--
ALTER TABLE ` + "`" + `person_type` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD UNIQUE KEY ` + "`" + `name` + "`" + ` (` + "`" + `name` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `person_with_lock` + "`" + `
--
ALTER TABLE ` + "`" + `person_with_lock` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `project` + "`" + `
--
ALTER TABLE ` + "`" + `project` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD UNIQUE KEY ` + "`" + `num` + "`" + ` (` + "`" + `num` + "`" + `),
    ADD KEY ` + "`" + `IDX_project_1` + "`" + ` (` + "`" + `project_status_type_id` + "`" + `),
    ADD KEY ` + "`" + `IDX_project_2` + "`" + ` (` + "`" + `manager_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `project_status_type` + "`" + `
--
ALTER TABLE ` + "`" + `project_status_type` + "`" + `
    ADD PRIMARY KEY (` + "`" + `id` + "`" + `),
    ADD UNIQUE KEY ` + "`" + `IDX_projectstatustype_1` + "`" + ` (` + "`" + `name` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `related_project_assn` + "`" + `
--
ALTER TABLE ` + "`" + `related_project_assn` + "`" + `
    ADD PRIMARY KEY (` + "`" + `parent_id` + "`" + `,` + "`" + `child_id` + "`" + `),
    ADD KEY ` + "`" + `IDX_relatedprojectassn_2` + "`" + ` (` + "`" + `child_id` + "`" + `);</p>
<p>--
-- Indexes for table ` + "`" + `team_member_project_assn` + "`" + `
--
ALTER TABLE ` + "`" + `team_member_project_assn` + "`" + `
    ADD PRIMARY KEY (` + "`" + `team_member_id` + "`" + `,` + "`" + `project_id` + "`" + `) USING BTREE,
    ADD KEY ` + "`" + `IDX_teammemberprojectassn_2` + "`" + ` (` + "`" + `project_id` + "`" + `);</p>
<p>--
-- AUTO_INCREMENT for dumped tables
--</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `address` + "`" + `
--
ALTER TABLE ` + "`" + `address` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=141;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `employee_info` + "`" + `
--
ALTER TABLE ` + "`" + `employee_info` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `login` + "`" + `
--
ALTER TABLE ` + "`" + `login` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `milestone` + "`" + `
--
ALTER TABLE ` + "`" + `milestone` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `person` + "`" + `
--
ALTER TABLE ` + "`" + `person` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=168;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `person_type` + "`" + `
--
ALTER TABLE ` + "`" + `person_type` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `person_with_lock` + "`" + `
--
ALTER TABLE ` + "`" + `person_with_lock` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `project` + "`" + `
--
ALTER TABLE ` + "`" + `project` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;</p>
<p>--
-- AUTO_INCREMENT for table ` + "`" + `project_status_type` + "`" + `
--
ALTER TABLE ` + "`" + `project_status_type` + "`" + `
    MODIFY ` + "`" + `id` + "`" + ` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;</p>
<p>--
-- Constraints for dumped tables
--</p>
<p>--
-- Constraints for table ` + "`" + `address` + "`" + `
--
ALTER TABLE ` + "`" + `address` + "`" + `
    ADD CONSTRAINT ` + "`" + `person_address` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `) ON DELETE CASCADE ON UPDATE CASCADE;</p>
<p>--
-- Constraints for table ` + "`" + `employee_info` + "`" + `
--
ALTER TABLE ` + "`" + `employee_info` + "`" + `
    ADD CONSTRAINT ` + "`" + `employee_info_ibfk_1` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `) ON DELETE CASCADE ON UPDATE CASCADE;</p>
<p>--
-- Constraints for table ` + "`" + `login` + "`" + `
--
ALTER TABLE ` + "`" + `login` + "`" + `
    ADD CONSTRAINT ` + "`" + `person_login` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `) ON DELETE CASCADE ON UPDATE CASCADE;</p>
<p>--
-- Constraints for table ` + "`" + `milestone` + "`" + `
--
ALTER TABLE ` + "`" + `milestone` + "`" + `
    ADD CONSTRAINT ` + "`" + `project_milestone` + "`" + ` FOREIGN KEY (` + "`" + `project_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `) ON DELETE CASCADE;</p>
<p>--
-- Constraints for table ` + "`" + `person_persontype_assn` + "`" + `
--
ALTER TABLE ` + "`" + `person_persontype_assn` + "`" + `
    ADD CONSTRAINT ` + "`" + `person_persontype_assn_1` + "`" + ` FOREIGN KEY (` + "`" + `person_type_id` + "`" + `) REFERENCES ` + "`" + `person_type` + "`" + ` (` + "`" + `id` + "`" + `),
    ADD CONSTRAINT ` + "`" + `person_persontype_assn_2` + "`" + ` FOREIGN KEY (` + "`" + `person_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `);</p>
<p>--
-- Constraints for table ` + "`" + `project` + "`" + `
--
ALTER TABLE ` + "`" + `project` + "`" + `
    ADD CONSTRAINT ` + "`" + `person_project` + "`" + ` FOREIGN KEY (` + "`" + `manager_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `),
    ADD CONSTRAINT ` + "`" + `project_status_type_project` + "`" + ` FOREIGN KEY (` + "`" + `project_status_type_id` + "`" + `) REFERENCES ` + "`" + `project_status_type` + "`" + ` (` + "`" + `id` + "`" + `);</p>
<p>--
-- Constraints for table ` + "`" + `related_project_assn` + "`" + `
--
ALTER TABLE ` + "`" + `related_project_assn` + "`" + `
    ADD CONSTRAINT ` + "`" + `related_project_assn_1` + "`" + ` FOREIGN KEY (` + "`" + `parent_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `),
    ADD CONSTRAINT ` + "`" + `related_project_assn_2` + "`" + ` FOREIGN KEY (` + "`" + `child_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `);</p>
<p>--
-- Constraints for table ` + "`" + `team_member_project_assn` + "`" + `
--
ALTER TABLE ` + "`" + `team_member_project_assn` + "`" + `
    ADD CONSTRAINT ` + "`" + `person_team_member_project_assn` + "`" + ` FOREIGN KEY (` + "`" + `team_member_id` + "`" + `) REFERENCES ` + "`" + `person` + "`" + ` (` + "`" + `id` + "`" + `) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT ` + "`" + `project_team_member_project_assn` + "`" + ` FOREIGN KEY (` + "`" + `project_id` + "`" + `) REFERENCES ` + "`" + `project` + "`" + ` (` + "`" + `id` + "`" + `) ON DELETE CASCADE ON UPDATE CASCADE;
SET FOREIGN_KEY_CHECKS=1;
</p>
`)

			buf.WriteString(`
</code>
`)

			buf.WriteString(`
</body>
</html>
`)

			return

		})
}
