#!/usr/bin/env node

const fs = require('fs');
const fastify = require('fastify')({
    logger: true
})

fastify.get('/v1/api/hello', async (request, reply) => {
    reply.code(200).send({ response: 'hello without api-key', success:true, data:'fastify node' })
})

fastify.get('/v1/api/healthy', async (request, reply) => {
    reply.code(200).send({ response: 'healthy', success:true, data:'fastify node' })
})

fastify.get('/v1/api/add', {
    schema: {
        querystring: {
            type: 'object',
            required: ['x', 'y'],
            properties: {
                x: { type: 'integer' },
                y: { type: 'integer' }
            }
        },
        response: {
            200: {
                type: 'object',
                properties: {
                    result: { type: 'integer' }
                }
            }
        }
    }
}, async (request, reply) => {
    const { x, y } = request.query;
    const result = x + y;
    reply.header('Content-Security-Policy',
        "default-src 'self'; font-src 'self'; img-src 'self'; script-src 'self'; style-src 'self'; frame-src 'self'")
    return { result };
});

fastify.listen(process.env.PORT || 3001, '::', function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
    fastify.log.info('We are up')
});

console.log('Node.js web server at port 3001 is running..');
