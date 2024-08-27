@description('The name of the AKS cluster')
param aksClusterName string = 'myAksCluster'

@description('The location for all resources.')
param location string = resourceGroup().location

@description('The Kubernetes version for the AKS cluster.')
param kubernetesVersion string = '1.24.6'

@description('The DNS prefix for the AKS cluster.')
param dnsPrefix string = 'myaksdns'

@description('The ID of the principal that will have Contributor role on the AKS cluster.')
param principalId string

@description('The client ID of the service principal used by the AKS cluster for authentication to Azure APIs')
param clientId string

@description('The client secret of the service principal used by the AKS cluster for authentication to Azure APIs')
param clientKey string

@description('The name of the storage account.')
param storageAccountName string = 'mystorageaccount'

@description('The name of the storage container.')
param containerName string = 'mycontainer'

resource aksCluster 'Microsoft.ContainerService/managedClusters@2023-01-01' = {
  name: aksClusterName
  location: location
  properties: {
    kubernetesVersion: kubernetesVersion
    dnsPrefix: dnsPrefix
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
  name: storageAccountName
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
  name: containerName
  parent: blobService
  properties: {
    publicAccess: 'None'
  }
}
