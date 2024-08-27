@description('The location for all resources.')
param location string = resourceGroup().location

@description('The name of the AKS cluster')
param aksClusterName string

@description('The client ID of the service principal used by the AKS cluster for authentication to Azure APIs')
param clientId string

@description('The client secret of the service principal used by the AKS cluster for authentication to Azure APIs')
param clientKey string

@description('The principal ID of the service principal used by the AKS cluster for RBAC')
param principalId string

resource aksCluster 'Microsoft.ContainerService/managedClusters@2023-01-01' = {
  name: aksClusterName
  location: location
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

resource roleAssignment 'Microsoft.Authorization/roleAssignments@2020-04-01-preview' = {
  name: guid(subscription().id, principalId, 'Contributor')
  scope: resourceGroup()
  properties: {
    roleDefinitionId: resourceId('Microsoft.Authorization/roleDefinitions', 'b24988ac-6180-42a0-ab88-20f7382dd24c')
    principalId: principalId
  }
}

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
  name: 'dmclusterblob'
  parent: storageAccount
}

resource storageContainer 'Microsoft.Storage/storageAccounts/blobServices/containers@2021-09-01' = {
  name: 'dmclustercontainer'
  parent: blobService
  properties: {
    publicAccess: 'None'
  }
}
