export const MESSAGES = {
  NEW_BUILDERS: "NEW_BUILDERS",
  START_BUILD: "START_BUILD",
  FINISHED_BUILD: "FINISHED_BUILD",
  FAIL_BUILD: "FAIL_BUILD",
  FINISH_ALL_BUILDERS: "FINISH_ALL_BUILDERS",
  START_GATHER: "START_GATHER",
  FINISHED_GATHER: "FINISHED_GATHER",
  FINISH_ALL_GATHERERS: "FINISH_ALL_GATHERERS",
  NEW_RESOURCES: "NEW_RESOURCES",
  NEW_GATHERERS: "NEW_GATHERERS",
  TOTAL_TIME: "TOTAL_TIME",
}

const getAction = (e) => {
  const [message, ...data] = e.split(' ')
  if(!MESSAGES[message]) return { type: 'default' }
  return { type: MESSAGES[message], data }
}

export default getAction
