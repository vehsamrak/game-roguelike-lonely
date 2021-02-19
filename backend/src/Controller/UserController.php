<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class UserController extends AbstractController
{
    /**
     * @Route("/user", name="users")
     */
    public function users(): Response
    {
        return $this->json([]);
    }
    /**
     * @Route("/user/{id}", name="user")
     */
    public function user($id): Response
    {
        return $this->json([
            'id' => $id
        ]);
    }
}
