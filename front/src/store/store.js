import { createStore } from 'redux'
import { MESSAGES } from './actions'

const startGather = (state, data) => {
  const [ material ] = data
  return {
    ...state,
    gatherers: {
      ...state.gatherers,
      [material]: state.gatherers[material] + 1
    },
    info: {
      ...state.info,
      jobsInProgress: state.info.jobsInProgress + 1
    }
  };
}

const finishGather = (state, data) => {
  const [ material, amount ] = data
  return {
    ...state,
    gatherers: {
      ...state.gatherers,
      [material]: state.gatherers[material] - 1
    },
    warehouse: {
      ...state.warehouse,
      [material]: state.warehouse[material] + Number(amount)
    },
    info: {
      ...state.info,
      jobsFinished: state.info.jobsFinished + 1,
      jobsInProgress: state.info.jobsInProgress - 1
    }
  };
}

const startBuild = (state, data) => {
  const [ material ] = data
  return {
    ...state,
    workers: {
      ...state.workers,
      [material]: state.workers[material] + 1
    },
    info: {
      ...state.info,
      jobsInProgress: state.info.jobsInProgress + 1
    }
  };
}

const finishBuild = (state, data) => {
  const [ weapon ] = data
  const materials = weapon === "sword" ?
    {
      stone: state.warehouse.stone - 20,
      gold: state.warehouse.gold - 10,
    } :
    {
      stone: state.warehouse.stone - 20,
      wood: state.warehouse.wood - 10,
    }
  return {
    ...state,
    workers: {
      ...state.workers,
      [weapon]: state.workers[weapon] - 1
    },
    warehouse: {
      ...state.warehouse,
      [weapon]: state.warehouse[weapon] + 1,
      ...materials
    },
    info: {
      ...state.info,
      jobsFinished: state.info.jobsFinished + 1,
      jobsInProgress: state.info.jobsInProgress - 1
    }
  };
}

const failBuild = (state, data) => {
  const [ material ] = data
  return {
    ...state,
    workers: {
      ...state.workers,
      [material]: state.workers[material] - 1
    },
    info: {
      ...state.info,
      jobsFinished: state.info.jobsFinished + 1,
      jobsInProgress: state.info.jobsInProgress - 1
    }
  };
}

const addWorkers = (state, data) => {
  const [ amount ] = data
  return {
    ...state,
    info: {
      ...state.info,
      totalWorkers: state.info.totalWorkers + Number(amount)
    }
  }
}

const addGatherers = (state, data) => {
  const [ amount ] = data
  return {
    ...state,
    info: {
      ...state.info,
      totalGatherers: state.info.totalGatherers + Number(amount)
    }
  }
}

const initialState = {
  gatherers: {
    wood: 0,
    stone: 0,
    gold: 0,
  },
  workers: {
    sword: 0,
    shield: 0,
  },
  warehouse: {
    wood: 0,
    stone: 0,
    gold: 0,
    sword: 0,
    shield: 0,
  },
  info: {
    jobsInProgress: 0,
    jobsFinished: 0,
    totalGatherers: 0,
    totalWorkers: 0,
  }
}

const reducerActions = {
  [MESSAGES.START_BUILD]: startBuild,
  [MESSAGES.FINISHED_BUILD]: finishBuild,
  [MESSAGES.FAIL_BUILD]: failBuild,
  [MESSAGES.START_GATHER]: startGather,
  [MESSAGES.FINISHED_GATHER]: finishGather,
  [MESSAGES.NEW_WORKERS]: addWorkers,
  [MESSAGES.NEW_GATHERERS]: addGatherers,
}

const reducer = (state = initialState, action) => {
  if(!reducerActions[action.type]) return state
  return reducerActions[action.type](state, action.data)
}

export const store = createStore(reducer)
