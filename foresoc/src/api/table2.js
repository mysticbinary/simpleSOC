import request from '@/utils/request'



export function findAll(query) {
  return request({
    url: '/api/user/findAll',
    method: 'get',
    params: query
  })
}