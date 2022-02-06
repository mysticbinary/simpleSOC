import request from '@/utils/request'

export function findAll() {
  return request({
    url: '/api/user/findAll',
    method: 'get',
    data
  })

}