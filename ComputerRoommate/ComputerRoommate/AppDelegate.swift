//
//  AppDelegate.swift
//  ComputerRoommate
//
//  Created by Dan Kurtz on 1/22/16.
//  Copyright © 2016 Dan Kurtz. All rights reserved.
//

import Cocoa
import Foundation

@NSApplicationMain
class AppDelegate: NSObject, NSApplicationDelegate {

    @IBOutlet weak var statusMenu: NSMenu?
    var statusItem: NSStatusItem? = nil;

    @IBAction func statusMenuItemSelected(sender: NSMenuItem) {
        let clipath = NSBundle.mainBundle().pathForResource("roommate-cli", ofType: nil)
        let icopath = NSBundle.mainBundle().pathForResource("icon", ofType: "icns")
        
        if sender.title == "Knock on door" {
            // Drop audio
            let task = NSTask()
            task.launchPath = clipath
            task.arguments = ["volume", "15"]
            task.launch()
        } else if sender.title == "Bang on door" {
            let task = NSTask()
            task.launchPath = clipath
            task.arguments = ["volume", "5"]
            task.launch()
        } else if sender.title == "Ask for rent" {
            let task = NSTask()
            task.launchPath = clipath
            task.arguments = ["alert", "yeah i'll totally get that to you soon ✌️✌️✌️", "Right, about that...", icopath!, "Wait", "Wait"]
            task.launch()
        }
    }
    
    func applicationDidFinishLaunching(aNotification: NSNotification) {
        // Insert code here to initialize your application
    }

    func applicationWillTerminate(aNotification: NSNotification) {
        // Insert code here to tear down your application
        let task = NSTask()
        task.launchPath = "/usr/bin/killall"
        task.arguments = ["roommate-cli"]
        task.launch()
    }
    
    override func awakeFromNib() {
        activateStatusMenu()
    }
    
    func activateStatusMenu() {
        let statusBar = NSStatusBar.systemStatusBar()
        
        statusItem = statusBar.statusItemWithLength(NSVariableStatusItemLength)
        statusItem!.title = "👩"
        statusItem!.menu = self.statusMenu
        NSApplication.sharedApplication().keyWindow?.backgroundColor = NSColor.whiteColor()
    }
    
}

