import request from '@/utils/request'
import data from '@/views/pdf/content'

export function findAll() {
  return request({
    url: '/api/user/findAll',
    method: 'get',
    data: {
      ...query,
      name,
    }
  })
  
}