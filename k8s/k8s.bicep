@description('The location for all resources.')
param location string = resourceGroup().location

@description('The name of the AKS cluster')
param aksClusterName string

@description('The client ID of the service principal used by the AKS cluster for authentication to Azure APIs')
param clientId string

@description('The client secret of the service principal used by the AKS cluster for authentication to Azure APIs')
param clientKey string

@description('Name of the Virtual Network')
param vnetName string = 'dm-vnet'

@description('Name of the Application Gateway Subnet')
param appGwSubnetName string = 'appGwSubnet'

@description('Address prefix for the Virtual Network')
param vnetAddressPrefix string = '10.1.0.0/16'

@description('Address prefix for the Application Gateway Subnet')
param appGwSubnetPrefix string = '10.1.1.0/24'

@description('Name of the Application Gateway')
param appGwName string = 'dm-app-gateway'

resource storageAccount 'Microsoft.Storage/storageAccounts@2021-04-01' = {
  name: 'dmclusterstorage'
  location: location
  sku: {
    name: 'Standard_LRS'
  }
  kind: 'StorageV2'
  properties: {}
}

resource blobService 'Microsoft.Storage/storageAccounts/blobServices@2021-04-01' = {
  name: 'default'
  parent: storageAccount
}

resource storageContainer 'Microsoft.Storage/storageAccounts/blobServices/containers@2021-09-01' = {
  name: 'dmclustercontainer'
  parent: blobService
  properties: {
    publicAccess: 'None'
  }
}

resource aksCluster 'Microsoft.ContainerService/managedClusters@2023-01-01' = {
  name: aksClusterName
  location: location
  dependsOn: [
    storageContainer
  ]
  properties: {
    kubernetesVersion: '1.29.0'
    dnsPrefix: 'dmc'
    agentPoolProfiles: [
      {
        name: 'nodepool1'
        count: 1
        vmSize: 'Standard_B2s'
        osType: 'Linux'
        mode: 'System'
      }
      {
        name: 'nodepool2'
        count: 1
        vmSize: 'Standard_B2s'
        osType: 'Linux'
        mode: 'User'
      }
    ]
    servicePrincipalProfile: {
      clientId: clientId
      secret: clientKey
    }
    networkProfile: {
      networkPlugin: 'azure'
      serviceCidr: '10.0.0.0/16'
      dnsServiceIP: '10.0.0.10'
      podCidr: '10.244.0.0/16'
    }
  }
}

resource vnet 'Microsoft.Network/virtualNetworks@2021-05-01' = {
  name: vnetName
  location: location
  properties: {
    addressSpace: {
      addressPrefixes: [
        vnetAddressPrefix
      ]
    }
    subnets: [
      {
        name: appGwSubnetName
        properties: {
          addressPrefix: appGwSubnetPrefix
        }
      }
    ]
  }
}

resource publicIPAddress 'Microsoft.Network/publicIPAddresses@2021-05-01' = {
  name: 'dm-app-gateway-pip'
  location: location
  properties: {
    publicIPAllocationMethod: 'Dynamic'
  }
}

resource appGw 'Microsoft.Network/applicationGateways@2021-05-01' = {
  name: appGwName
  location: location
  properties: {
    sku: {
      name: 'Standard_v2'
      tier: 'Standard_v2'
      capacity: 2
    }
    gatewayIPConfigurations: [
      {
        name: 'appGwIpConfig'
        properties: {
          subnet: {
            id: resourceId('Microsoft.Network/virtualNetworks/subnets', vnet.name, appGwSubnetName)
          }
        }
      }
    ]
    frontendIPConfigurations: [
      {
        name: 'appGatewayFrontendIP'
        properties: {
          publicIPAddress: {
            id: publicIPAddress.id
          }
        }
      }
    ]
    frontendPorts: [
      {
        name: 'appGatewayFrontendPort'
        properties: {
          port: 443
        }
      }
    ]
    backendAddressPools: [
      {
        name: 'appGatewayBackendPool'
      }
    ]
    backendHttpSettingsCollection: [
      {
        name: 'appGatewayBackendHttpSettings'
        properties: {
          cookieBasedAffinity: 'Disabled'
          port: 80
          protocol: 'Http'
        }
      }
    ]
    httpListeners: [
      {
        name: 'appGatewayHttpListener'
        properties: {
          frontendIPConfiguration: {
            id: resourceId('Microsoft.Network/applicationGateways/frontendIPConfigurations', 'appGatewayFrontendIP')
          }
          frontendPort: {
            id: resourceId('Microsoft.Network/applicationGateways/frontendPorts', 'appGatewayFrontendPort')
          }
          protocol: 'Https'
          sslCertificate: {
            id: resourceId('Microsoft.Network/applicationGateways/sslCertificates', 'appGatewaySslCert')
          }
        }
      }
    ]
    urlPathMaps: [
      {
        name: 'appGatewayUrlPathMap'
        properties: {
          defaultBackendAddressPool: {
            id: resourceId('Microsoft.Network/applicationGateways/backendAddressPools', 'appGatewayBackendPool')
          }
          defaultBackendHttpSettings: {
            id: resourceId(
              'Microsoft.Network/applicationGateways/backendHttpSettingsCollection',
              'appGatewayBackendHttpSettings'
            )
          }
        }
      }
    ]
  }
}
