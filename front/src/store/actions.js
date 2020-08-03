export const MESSAGES = {
  START_BUILD: "START_BUILD",
  FINISHED_BUILD: "FINISHED_BUILD",
  FAIL_BUILD: "FAIL_BUILD",
  START_GATHER: "START_GATHER",
  FINISHED_GATHER: "FINISHED_GATHER",
  NEW_WORKERS: "NEW_WORKERS",
  NEW_GATHERERS: "NEW_GATHERERS",
}

const getAction = (e) => {
  const [message, ...data] = e.split(' ')
  if(!MESSAGES[message]) return { type: 'default' }
  return { type: MESSAGES[message], data }
}

export default getAction
