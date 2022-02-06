import request from '@/utils/request'

export function serviceList(query) {
    return request({
        url: '/api/service/list',
        method: 'get',
        params: query
    })
}