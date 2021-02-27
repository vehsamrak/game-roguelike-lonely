<?php

namespace App\Tests\Controller;

use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

class UserControllerTest extends WebTestCase
{
    public function testShowPost()
    {
        $client = static::createClient();

        $client->request('GET', '/user');

        $this->assertEquals(200, $client->getResponse()->getStatusCode());
    }

    /**
     * @test
     */
    public function environmentMustBeTest(): void
    {
       static::bootKernel();

       $environment = static::$kernel->getEnvironment();

       $this->assertSame('test', $environment);
    }
}
