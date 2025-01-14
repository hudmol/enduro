/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    CollectionBulkNotAvailableResponseBody,
    CollectionBulkNotAvailableResponseBodyFromJSON,
    CollectionBulkNotAvailableResponseBodyToJSON,
    CollectionBulkNotValidResponseBody,
    CollectionBulkNotValidResponseBodyFromJSON,
    CollectionBulkNotValidResponseBodyToJSON,
    CollectionBulkRequestBody,
    CollectionBulkRequestBodyFromJSON,
    CollectionBulkRequestBodyToJSON,
    CollectionBulkResponseBody,
    CollectionBulkResponseBodyFromJSON,
    CollectionBulkResponseBodyToJSON,
    CollectionBulkStatusResponseBody,
    CollectionBulkStatusResponseBodyFromJSON,
    CollectionBulkStatusResponseBodyToJSON,
    CollectionCancelNotFoundResponseBody,
    CollectionCancelNotFoundResponseBodyFromJSON,
    CollectionCancelNotFoundResponseBodyToJSON,
    CollectionCancelNotRunningResponseBody,
    CollectionCancelNotRunningResponseBodyFromJSON,
    CollectionCancelNotRunningResponseBodyToJSON,
    CollectionDecideNotFoundResponseBody,
    CollectionDecideNotFoundResponseBodyFromJSON,
    CollectionDecideNotFoundResponseBodyToJSON,
    CollectionDecideNotValidResponseBody,
    CollectionDecideNotValidResponseBodyFromJSON,
    CollectionDecideNotValidResponseBodyToJSON,
    CollectionDeleteNotFoundResponseBody,
    CollectionDeleteNotFoundResponseBodyFromJSON,
    CollectionDeleteNotFoundResponseBodyToJSON,
    CollectionDownloadNotFoundResponseBody,
    CollectionDownloadNotFoundResponseBodyFromJSON,
    CollectionDownloadNotFoundResponseBodyToJSON,
    CollectionListResponseBody,
    CollectionListResponseBodyFromJSON,
    CollectionListResponseBodyToJSON,
    CollectionMonitorResponseBody,
    CollectionMonitorResponseBodyFromJSON,
    CollectionMonitorResponseBodyToJSON,
    CollectionRetryNotFoundResponseBody,
    CollectionRetryNotFoundResponseBodyFromJSON,
    CollectionRetryNotFoundResponseBodyToJSON,
    CollectionRetryNotRunningResponseBody,
    CollectionRetryNotRunningResponseBodyFromJSON,
    CollectionRetryNotRunningResponseBodyToJSON,
    CollectionShowNotFoundResponseBody,
    CollectionShowNotFoundResponseBodyFromJSON,
    CollectionShowNotFoundResponseBodyToJSON,
    CollectionShowResponseBody,
    CollectionShowResponseBodyFromJSON,
    CollectionShowResponseBodyToJSON,
    CollectionWorkflowNotFoundResponseBody,
    CollectionWorkflowNotFoundResponseBodyFromJSON,
    CollectionWorkflowNotFoundResponseBodyToJSON,
    CollectionWorkflowResponseBody,
    CollectionWorkflowResponseBodyFromJSON,
    CollectionWorkflowResponseBodyToJSON,
    InlineObject,
    InlineObjectFromJSON,
    InlineObjectToJSON,
} from '../models';

export interface CollectionBulkRequest {
    bulkRequestBody: CollectionBulkRequestBody;
}

export interface CollectionCancelRequest {
    id: number;
}

export interface CollectionDecideRequest {
    id: number;
    object: InlineObject;
}

export interface CollectionDeleteRequest {
    id: number;
}

export interface CollectionDownloadRequest {
    id: number;
}

export interface CollectionListRequest {
    name?: string;
    originalId?: string;
    transferId?: string;
    aipId?: string;
    pipelineId?: string;
    earliestCreatedTime?: Date;
    latestCreatedTime?: Date;
    status?: CollectionListStatusEnum;
    cursor?: string;
}

export interface CollectionRetryRequest {
    id: number;
}

export interface CollectionShowRequest {
    id: number;
}

export interface CollectionWorkflowRequest {
    id: number;
}

/**
 * no description
 */
export class CollectionApi extends runtime.BaseAPI {

    /**
     * Bulk operations (retry, cancel...).
     * bulk collection
     */
    async collectionBulkRaw(requestParameters: CollectionBulkRequest): Promise<runtime.ApiResponse<CollectionBulkResponseBody>> {
        if (requestParameters.bulkRequestBody === null || requestParameters.bulkRequestBody === undefined) {
            throw new runtime.RequiredError('bulkRequestBody','Required parameter requestParameters.bulkRequestBody was null or undefined when calling collectionBulk.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/collection/bulk`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CollectionBulkRequestBodyToJSON(requestParameters.bulkRequestBody),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionBulkResponseBodyFromJSON(jsonValue));
    }

    /**
     * Bulk operations (retry, cancel...).
     * bulk collection
     */
    async collectionBulk(requestParameters: CollectionBulkRequest): Promise<CollectionBulkResponseBody> {
        const response = await this.collectionBulkRaw(requestParameters);
        return await response.value();
    }

    /**
     * Retrieve status of current bulk operation.
     * bulk_status collection
     */
    async collectionBulkStatusRaw(): Promise<runtime.ApiResponse<CollectionBulkStatusResponseBody>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/bulk`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionBulkStatusResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve status of current bulk operation.
     * bulk_status collection
     */
    async collectionBulkStatus(): Promise<CollectionBulkStatusResponseBody> {
        const response = await this.collectionBulkStatusRaw();
        return await response.value();
    }

    /**
     * Cancel collection processing by ID
     * cancel collection
     */
    async collectionCancelRaw(requestParameters: CollectionCancelRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionCancel.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/cancel`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Cancel collection processing by ID
     * cancel collection
     */
    async collectionCancel(requestParameters: CollectionCancelRequest): Promise<void> {
        await this.collectionCancelRaw(requestParameters);
    }

    /**
     * Make decision for a pending collection by ID
     * decide collection
     */
    async collectionDecideRaw(requestParameters: CollectionDecideRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionDecide.');
        }

        if (requestParameters.object === null || requestParameters.object === undefined) {
            throw new runtime.RequiredError('object','Required parameter requestParameters.object was null or undefined when calling collectionDecide.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/collection/{id}/decision`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: InlineObjectToJSON(requestParameters.object),
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Make decision for a pending collection by ID
     * decide collection
     */
    async collectionDecide(requestParameters: CollectionDecideRequest): Promise<void> {
        await this.collectionDecideRaw(requestParameters);
    }

    /**
     * Delete collection by ID
     * delete collection
     */
    async collectionDeleteRaw(requestParameters: CollectionDeleteRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionDelete.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete collection by ID
     * delete collection
     */
    async collectionDelete(requestParameters: CollectionDeleteRequest): Promise<void> {
        await this.collectionDeleteRaw(requestParameters);
    }

    /**
     * Download collection by ID
     * download collection
     */
    async collectionDownloadRaw(requestParameters: CollectionDownloadRequest): Promise<runtime.ApiResponse<string>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionDownload.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/download`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.TextApiResponse(response) as any;
    }

    /**
     * Download collection by ID
     * download collection
     */
    async collectionDownload(requestParameters: CollectionDownloadRequest): Promise<string> {
        const response = await this.collectionDownloadRaw(requestParameters);
        return await response.value();
    }

    /**
     * List all stored collections
     * list collection
     */
    async collectionListRaw(requestParameters: CollectionListRequest): Promise<runtime.ApiResponse<CollectionListResponseBody>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.originalId !== undefined) {
            queryParameters['original_id'] = requestParameters.originalId;
        }

        if (requestParameters.transferId !== undefined) {
            queryParameters['transfer_id'] = requestParameters.transferId;
        }

        if (requestParameters.aipId !== undefined) {
            queryParameters['aip_id'] = requestParameters.aipId;
        }

        if (requestParameters.pipelineId !== undefined) {
            queryParameters['pipeline_id'] = requestParameters.pipelineId;
        }

        if (requestParameters.earliestCreatedTime !== undefined) {
            queryParameters['earliest_created_time'] = (requestParameters.earliestCreatedTime as any).toISOString();
        }

        if (requestParameters.latestCreatedTime !== undefined) {
            queryParameters['latest_created_time'] = (requestParameters.latestCreatedTime as any).toISOString();
        }

        if (requestParameters.status !== undefined) {
            queryParameters['status'] = requestParameters.status;
        }

        if (requestParameters.cursor !== undefined) {
            queryParameters['cursor'] = requestParameters.cursor;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionListResponseBodyFromJSON(jsonValue));
    }

    /**
     * List all stored collections
     * list collection
     */
    async collectionList(requestParameters: CollectionListRequest): Promise<CollectionListResponseBody> {
        const response = await this.collectionListRaw(requestParameters);
        return await response.value();
    }

    /**
     * monitor collection
     */
    async collectionMonitorRaw(): Promise<runtime.ApiResponse<void>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/monitor`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * monitor collection
     */
    async collectionMonitor(): Promise<void> {
        await this.collectionMonitorRaw();
    }

    /**
     * Retry collection processing by ID
     * retry collection
     */
    async collectionRetryRaw(requestParameters: CollectionRetryRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionRetry.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/retry`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Retry collection processing by ID
     * retry collection
     */
    async collectionRetry(requestParameters: CollectionRetryRequest): Promise<void> {
        await this.collectionRetryRaw(requestParameters);
    }

    /**
     * Show collection by ID
     * show collection
     */
    async collectionShowRaw(requestParameters: CollectionShowRequest): Promise<runtime.ApiResponse<CollectionShowResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionShow.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionShowResponseBodyFromJSON(jsonValue));
    }

    /**
     * Show collection by ID
     * show collection
     */
    async collectionShow(requestParameters: CollectionShowRequest): Promise<CollectionShowResponseBody> {
        const response = await this.collectionShowRaw(requestParameters);
        return await response.value();
    }

    /**
     * Retrieve workflow status by ID
     * workflow collection
     */
    async collectionWorkflowRaw(requestParameters: CollectionWorkflowRequest): Promise<runtime.ApiResponse<CollectionWorkflowResponseBody>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling collectionWorkflow.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/collection/{id}/workflow`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => CollectionWorkflowResponseBodyFromJSON(jsonValue));
    }

    /**
     * Retrieve workflow status by ID
     * workflow collection
     */
    async collectionWorkflow(requestParameters: CollectionWorkflowRequest): Promise<CollectionWorkflowResponseBody> {
        const response = await this.collectionWorkflowRaw(requestParameters);
        return await response.value();
    }

}

/**
    * @export
    * @enum {string}
    */
export enum CollectionListStatusEnum {
    New = 'new',
    InProgress = 'in progress',
    Done = 'done',
    Error = 'error',
    Unknown = 'unknown',
    Queued = 'queued',
    Pending = 'pending',
    Abandoned = 'abandoned'
}
