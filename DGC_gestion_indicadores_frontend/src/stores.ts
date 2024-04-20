import { writable } from "svelte/store";

export let hasTeacherDeleted = writable(false)
export let updateTeacherAction = writable({ status: false, teacherID: -1 })