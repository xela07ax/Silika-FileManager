export function someMutation (/* state */) {
}
export const setLoading = (state, opened) => { // Открыто или закрыто это окно
  state.loading = opened
}
export const setList = (state, data) => { // Открыто или закрыто это окно
  state.list = data
}
export const setDised = (state, data) => { // Открыто или закрыто это окно
  //console.log(data)
  const clone = JSON.parse(JSON.stringify(state.dised));
  clone[data.id] = data.val
  state.dised = clone
}
