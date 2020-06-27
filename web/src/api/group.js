import request from '@/utils/request'

/**
 * @return {Promise}
 */
export function getList({ page, rows }) {
  return request({
    url: '/group/getList',
    method: 'get',
    params: { page, rows }
  })
}

/**
 * @return {Promise}
 */
export function getTotal() {
  return request({
    url: '/group/getTotal',
    method: 'get',
    params: { }
  })
}

/**
 * @return {Promise}
 */
export function getOption() {
  return request({
    url: '/group/getOption',
    method: 'get'
  })
}

/**
 * @return {Promise}
 */
export function getDeployOption() {
  return request({
    url: '/group/getDeployOption',
    method: 'get'
  })
}

export function add(data) {
  return request({
    url: '/group/add',
    method: 'post',
    data
  })
}

export function edit(data) {
  return request({
    url: '/group/edit',
    method: 'post',
    data
  })
}

export function remove(id) {
  return request({
    url: '/group/remove',
    method: 'delete',
    data: { id }
  })
}
