import request from '@/utils/request'

/**
 * @return {Promise}
 */
export function getList({ page, rows }) {
  return request({
    url: '/package/getList',
    method: 'get',
    params: { page, rows }
  })
}

/**
 * @return {Promise}
 */
export function getOption() {
  return request({
    url: '/package/getOption',
    method: 'get'
  })
}
