import request from '@/utils/request'

export function add(data) {
  return request({
    url: '/api/add',
    method: 'post',
    data
  })
}

export function del(id) {
  return request({
    url: '/api/del',
    method: 'get',
    params: { id }
  })
}

export function delIds(ids) {
  return request({
    url: '/api/del/ids',
    method: 'get',
    params: { ids }
  })
}

export function getList(params) {
  return request({
    url: '/api/list',
    method: 'get',
    params
  })
}

export function searchDel(params) {
  return request({
    url: '/api/search-del',
    method: 'get',
    params
  })
}
