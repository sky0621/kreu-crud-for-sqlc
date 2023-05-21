/*-------------------------------------------------------------------------
 *
 * pg_proc_d.h
 *    Macro definitions for pg_proc
 *
 * Portions Copyright (c) 1996-2022, PostgreSQL Global Development Group
 * Portions Copyright (c) 1994, Regents of the University of California
 *
 * NOTES
 *  ******************************
 *  *** DO NOT EDIT THIS FILE! ***
 *  ******************************
 *
 *  It has been GENERATED by src/backend/catalog/genbki.pl
 *
 *-------------------------------------------------------------------------
 */
#ifndef PG_PROC_D_H
#define PG_PROC_D_H

#define ProcedureRelationId 1255
#define ProcedureRelation_Rowtype_Id 81
#define ProcedureOidIndexId 2690
#define ProcedureNameArgsNspIndexId 2691

#define Anum_pg_proc_oid 1
#define Anum_pg_proc_proname 2
#define Anum_pg_proc_pronamespace 3
#define Anum_pg_proc_proowner 4
#define Anum_pg_proc_prolang 5
#define Anum_pg_proc_procost 6
#define Anum_pg_proc_prorows 7
#define Anum_pg_proc_provariadic 8
#define Anum_pg_proc_prosupport 9
#define Anum_pg_proc_prokind 10
#define Anum_pg_proc_prosecdef 11
#define Anum_pg_proc_proleakproof 12
#define Anum_pg_proc_proisstrict 13
#define Anum_pg_proc_proretset 14
#define Anum_pg_proc_provolatile 15
#define Anum_pg_proc_proparallel 16
#define Anum_pg_proc_pronargs 17
#define Anum_pg_proc_pronargdefaults 18
#define Anum_pg_proc_prorettype 19
#define Anum_pg_proc_proargtypes 20
#define Anum_pg_proc_proallargtypes 21
#define Anum_pg_proc_proargmodes 22
#define Anum_pg_proc_proargnames 23
#define Anum_pg_proc_proargdefaults 24
#define Anum_pg_proc_protrftypes 25
#define Anum_pg_proc_prosrc 26
#define Anum_pg_proc_probin 27
#define Anum_pg_proc_prosqlbody 28
#define Anum_pg_proc_proconfig 29
#define Anum_pg_proc_proacl 30

#define Natts_pg_proc 30


/*
 * Symbolic values for prokind column
 */
#define PROKIND_FUNCTION 'f'
#define PROKIND_AGGREGATE 'a'
#define PROKIND_WINDOW 'w'
#define PROKIND_PROCEDURE 'p'

/*
 * Symbolic values for provolatile column: these indicate whether the result
 * of a function is dependent *only* on the values of its explicit arguments,
 * or can change due to outside factors (such as parameter variables or
 * table contents).  NOTE: functions having side-effects, such as setval(),
 * must be labeled volatile to ensure they will not get optimized away,
 * even if the actual return value is not changeable.
 */
#define PROVOLATILE_IMMUTABLE	'i' /* never changes for given input */
#define PROVOLATILE_STABLE		's' /* does not change within a scan */
#define PROVOLATILE_VOLATILE	'v' /* can change even within a scan */

/*
 * Symbolic values for proparallel column: these indicate whether a function
 * can be safely be run in a parallel backend, during parallelism but
 * necessarily in the leader, or only in non-parallel mode.
 */
#define PROPARALLEL_SAFE		's' /* can run in worker or leader */
#define PROPARALLEL_RESTRICTED	'r' /* can run in parallel leader only */
#define PROPARALLEL_UNSAFE		'u' /* banned while in parallel mode */

/*
 * Symbolic values for proargmodes column.  Note that these must agree with
 * the FunctionParameterMode enum in parsenodes.h; we declare them here to
 * be accessible from either header.
 */
#define PROARGMODE_IN		'i'
#define PROARGMODE_OUT		'o'
#define PROARGMODE_INOUT	'b'
#define PROARGMODE_VARIADIC 'v'
#define PROARGMODE_TABLE	't'


#endif							/* PG_PROC_D_H */