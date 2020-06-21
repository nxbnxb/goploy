import request from '@/utils/request'

/**
 * @return {Promise}
 */
export function getList(pagination) {
  return request({
    url: '/crontab/getList',
    method: 'get',
    params: { ...pagination }
  })
}

export function add(data) {
  return request({
    url: '/crontab/add',
    method: 'post',
    data
  })
}

export function edit(data) {
  return request({
    url: '/crontab/edit',
    method: 'post',
    data
  })
}

export function remove(id) {
  return request({
    url: '/crontab/remove',
    method: 'delete',
    data: { id }
  })
}
