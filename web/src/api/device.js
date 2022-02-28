import service from '@/utils/request'

export const getDeviceList = (data) => {
  console.log(data)
  return service({
    url: '/device/getDeviceList',
    method: 'post',
    data: data
  })
}

export const updateDeviceHolder = (data) => {
  return service({
    url: '/device/updateDeviceHolder',
    method: 'post',
    data: data
  })
}
